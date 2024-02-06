package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// TestNewToken tests new a token create operation.
func TestNewToken(t *testing.T) {
	id := bson.NewObjectId()
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// TestValidateToken tests the validation of a token.
func TestValidateToken(t *testing.T) {
	// Generate a new ObjectId for testing purposes
	id := bson.NewObjectId()

	// Create a new token with the generated ObjectId
	token, err := NewToken(id.Hex())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the token and ensure no errors
	userId, err := ValidateToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, userId)
	assert.Equal(t, id.Hex(), userId)

	// Wait for the token to expire (e.g., 1 second) to test expiration
	time.Sleep(time.Second)

	// Generate an expired token for testing purposes
	tokenExpired := GetTokenExpired(id.Hex())

	// Attempt to validate the expired token and ensure it returns an error
	userId, err = ValidateToken(tokenExpired)
	assert.Error(t, err)
	assert.Equal(t, userId, "")
	assert.EqualError(t, err, "token has invalid claims: token is expired")

}

// GetTokenExpired tests the operation of receiving the token expire duration.
func GetTokenExpired(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Minute * 5 * -1).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecretKey)
	return tokenString
}
