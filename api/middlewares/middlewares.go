package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/api/util"
	"github.com/mtnmunuklu/bavul/security"
)

// LogRequests provides logging of incoming requests.
func LogRequests(c *fiber.Ctx) error {
	t := time.Now()
	err := c.Next()
	log.Printf(`{"proto": "%s", "method": "%s", "route": "%s%s", "request_time": "%v"}`, c.Protocol(), c.Method(), c.Hostname(), c.Path(), time.Since(t))
	return err
}

// Authenticate provides the authentication process middleware.
func Authenticate(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString, err := security.ExtractToken(c)
		if err != nil {
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("error on parse token:", err.Error())
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		if !token.Valid {
			log.Println("invalid token:", tokenString)
			util.WriteError(c, fiber.StatusUnauthorized, util.ErrUnauthorized)
			return nil
		}

		return next(c)
	}
}

// CORS provides Cross-Origin Resource Sharing middleware.
func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}
