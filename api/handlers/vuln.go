package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/api/util"
	"github.com/mtnmunuklu/bavul/pb"
)

// VulnHandlers is the interface of the vulnerability operation.
type VulnHandlers interface {
	AddCVE(c *fiber.Ctx) error
	SearchCVE(c *fiber.Ctx) error
	GetAllCVEs(c *fiber.Ctx) error
	DeleteCVE(c *fiber.Ctx) error
	UpdateCVE(c *fiber.Ctx) error
	FetchNVDFeeds(c *fiber.Ctx) error
}

// vulnHandlers provides a connection with vulnerability service over proto buffer.
type vulnHandlers struct {
	authSvcClient pb.AuthServiceClient
	vulnSvcClient pb.VulnServiceClient
}

// NewAuthHandlers creates a new VulnHandlers instance.
func NewVulnHandlers(authSvcClient pb.AuthServiceClient, vulnSvcClient pb.VulnServiceClient) VulnHandlers {
	return &vulnHandlers{authSvcClient: authSvcClient, vulnSvcClient: vulnSvcClient}
}

func (h *vulnHandlers) AddCVE(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	addCVERequest := new(pb.AddCVERequest)
	if err := c.BodyParser(addCVERequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	addedCVE, err := h.vulnSvcClient.AddCVE(c.Context(), addCVERequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, addedCVE)

}

func (h *vulnHandlers) GetAllCVEs(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	// Cache key creation
	cacheKey := "GetAllCVEs"

	// Get value from cache
	if cachedData, found := util.GetFromCache(cacheKey); found {
		return util.WriteAsJSON(c, http.StatusOK, cachedData)
	}

	stream, err := h.vulnSvcClient.GetAllCVEs(c.Context(), &pb.GetAllCVEsRequest{})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var getedCVEs []*pb.CVE
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		getedCVEs = append(getedCVEs, user)
	}

	// Add the obtained data to the cache
	util.SetToCache(cacheKey, getedCVEs, 5*time.Minute)

	return util.WriteAsJSON(c, http.StatusOK, getedCVEs)
}

func (h *vulnHandlers) DeleteCVE(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	cveId := c.Get("CveId")
	deleteCVERequest := &pb.DeleteCVERequest{CveId: cveId}

	deletedCVE, err := h.vulnSvcClient.DeleteCVE(c.Context(), deleteCVERequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, deletedCVE)
}

func (h *vulnHandlers) UpdateCVE(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	updateCVERequest := new(pb.UpdateCVERequest)
	if err := c.BodyParser(updateCVERequest); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	updatedCVE, err := h.vulnSvcClient.UpdateCVE(c.Context(), updateCVERequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, updatedCVE)
}

func (h *vulnHandlers) FetchNVDFeeds(c *fiber.Ctx) error {
	userId, err := util.GetUserIDFromToken(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	getedUserRole, err := h.authSvcClient.GetUserRole(c.Context(), &pb.GetUserRoleRequest{Id: userId})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	userIsAdmin := util.CheckUserIsAdmin(getedUserRole.Role)
	if !userIsAdmin {
		return util.WriteError(c, http.StatusUnauthorized, util.ErrUnauthorized)
	}

	apiKey := c.Get("ApiKey")

	// Cache key creation
	cacheKey := "FetchNVDFeeds:" + apiKey

	// Get value from cache
	if cachedData, found := util.GetFromCache(cacheKey); found {
		return util.WriteAsJSON(c, http.StatusOK, cachedData)
	}

	stream, err := h.vulnSvcClient.FetchNVDFeeds(c.Context(), &pb.FetchNVDFeedsRequest{ApiKey: apiKey})
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var fetchedNVDFeeds []*pb.CVE
	for {
		fetchedNVDFeed, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		fetchedNVDFeeds = append(fetchedNVDFeeds, fetchedNVDFeed)
	}

	// Add the obtained data to the cache
	util.SetToCache(cacheKey, fetchedNVDFeeds, 5*time.Minute)

	return util.WriteAsJSON(c, http.StatusOK, fetchedNVDFeeds)
}

func (h *vulnHandlers) SearchCVE(c *fiber.Ctx) error {
	// Cache key creation
	cacheKey := fmt.Sprintf("SearchCVE:%s:%s:%s:%s:%s:%s",
		c.Get("CveId"), c.Get("Severity"), c.Get("Product"),
		c.Get("Vendor"), c.Get("StartDate"), c.Get("EndDate"))

	// Get value from cache
	if cachedData, found := util.GetFromCache(cacheKey); found {
		return util.WriteAsJSON(c, http.StatusOK, cachedData)
	}

	searchCVEsRequest := &pb.SearchCVERequest{
		CveId:     c.Get("CveId"),
		Severity:  c.Get("Severity"),
		Product:   c.Get("Product"),
		Vendor:    c.Get("Vendor"),
		StartDate: c.Get("StartDate"),
		EndDate:   c.Get("EndDate"),
	}

	stream, err := h.vulnSvcClient.SearchCVE(c.Context(), searchCVEsRequest)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	var searchedCVEs []*pb.CVE
	for {
		searchedCVE, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return util.WriteError(c, http.StatusBadRequest, err)
		}

		searchedCVEs = append(searchedCVEs, searchedCVE)
	}

	// Add the obtained data to the cache
	util.SetToCache(cacheKey, searchedCVEs, 5*time.Minute)

	return util.WriteAsJSON(c, http.StatusOK, searchedCVEs)
}
