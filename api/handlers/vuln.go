package handlers

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/api/util"
	"github.com/mtnmunuklu/bavul/pb"
)

// VulnHandlers is the interface of the vulnerability operation.
type VulnHandlers interface {
	AddCVE(c *fiber.Ctx) error
	GetCVE(c *fiber.Ctx) error
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
func NewVulnHandlers(vulnSvcClient pb.VulnServiceClient) VulnHandlers {
	return &vulnHandlers{vulnSvcClient: vulnSvcClient}
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

	cve := new(pb.AddCVERequest)
	if err := c.BodyParser(cve); err != nil {
		return util.WriteError(c, http.StatusBadRequest, err)
	}

	addedCVE, err := h.vulnSvcClient.AddCVE(c.Context(), cve)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, addedCVE)

}

func (h *vulnHandlers) GetCVE(c *fiber.Ctx) error {
	cveId := c.Get("cveId")
	cve := &pb.GetCVERequest{CveId: cveId}

	getedCVE, err := h.vulnSvcClient.GetCVE(c.Context(), cve)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, getedCVE)
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

	cveId := c.Get("cveId")
	cve := &pb.DeleteCVERequest{CveId: cveId}

	deletedCVE, err := h.vulnSvcClient.DeleteCVE(c.Context(), cve)
	if err != nil {
		return util.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	return util.WriteAsJSON(c, http.StatusOK, deletedCVE)
}

func (h *vulnHandlers) UpdateCVE(c *fiber.Ctx) error {
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
	fetchNVDFeedsRequest := &pb.FetchNVDFeedsRequest{ApiKey: apiKey}

	stream, err := h.vulnSvcClient.FetchNVDFeeds(c.Context(), fetchNVDFeedsRequest)
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

	return util.WriteAsJSON(c, http.StatusOK, fetchedNVDFeeds)
}
