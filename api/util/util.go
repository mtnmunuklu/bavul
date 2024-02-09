package util

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mtnmunuklu/bavul/security"
	"github.com/mtnmunuklu/bicache"
)

// Contains error codes for API.
var (
	ErrEmptyBody    = errors.New("body can't be empty")
	ErrEmptyHeader  = errors.New("header can't be empty")
	ErrExistURL     = errors.New("URL already exists")
	ErrUnauthorized = errors.New("unauthorized operation")
)

// JError represents an error structure.
type JError struct {
	Error string `json:"error"`
}

// WriteAsJSON writes the response in JSON format.
func WriteAsJSON(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Set("Content-Type", "application/json")
	return c.Status(statusCode).JSON(data)
}

// WriteError writes the error response in JSON format.
func WriteError(c *fiber.Ctx, statusCode int, err error) error {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	return WriteAsJSON(c, statusCode, JError{Error: e})
}

// GetUserIDFromToken returns the user ID from the token.
func GetUserIDFromToken(c *fiber.Ctx) (string, error) {
	token, err := security.ExtractToken(c)
	if err != nil {
		return "", err
	}

	userID, err := security.ValidateToken(token)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// CheckUserIsAdmin checks if the user is an admin.
func CheckUserIsAdmin(role string) bool {
	return role == "admin"
}

// Create a BiCache instance
var myCache = bicache.NewBiCache(1000, 10*time.Minute)

func GetFromCache(key string) (interface{}, bool) {
	return myCache.Get(key)
}

func SetToCache(key string, value interface{}, expiration time.Duration) {
	myCache.Set(key, value, expiration)
}
