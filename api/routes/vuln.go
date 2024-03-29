package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/api/handlers"
)

// NewVulnRoutes provides the routing process for vulnerability.
func NewVulnRoutes(vulnHandlers handlers.VulnHandlers) []*Route {
	return []*Route{
		{
			Method: http.MethodPut,
			Path:   "/cve",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.AddCVE(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/cve",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.SearchCVE(c)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/cves",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.GetAllCVEs(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodDelete,
			Path:   "/cve",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.DeleteCVE(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodPost,
			Path:   "/cve",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.UpdateCVE(c)
			},
			AuthRequired: true,
		},
		{
			Method: http.MethodGet,
			Path:   "/nvd",
			Handler: func(c *fiber.Ctx) error {
				return vulnHandlers.FetchNVDFeeds(c)
			},
			AuthRequired: true,
		},
	}
}
