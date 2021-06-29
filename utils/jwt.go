package utils

import (
	"errors"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rajeshj3/jwt-auth/config"
	"github.com/rajeshj3/jwt-auth/database"
	"github.com/rajeshj3/jwt-auth/models"
	"github.com/rajeshj3/jwt-auth/responses"
)

// create JWT
func CreateJWT(user_id uint32) (string, error) {
	// calim using HS256
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user_id)),
		ExpiresAt: config.JWTExpireLife.Unix(),
	})

	// create token hashed with SECRET_KEY
	return claims.SignedString([]byte(config.SECRET_KEY))
}

// Prepare and SetCookie to header
func CreateJWTCookie(c *fiber.Ctx, id uint32) error {
	// Get Token
	token, err := CreateJWT(id)
	if err != nil {
		return responses.ErrorResponse(c, err)
	}

	// Prepare Cookie
	cookie := fiber.Cookie{
		Name:     config.JWTName,
		Value:    token,
		Expires:  config.JWTExpireLife,
		HTTPOnly: config.JWTHTTPOnly,
	}

	// SetCookie to header
	c.Cookie(&cookie)
	return nil
}

func CreateBlankJWTCookie(c *fiber.Ctx) error {
	// Prepare Blank Cookie
	cookie := fiber.Cookie{
		Name:     config.JWTName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: config.JWTHTTPOnly,
	}

	// SetCookie to header
	c.Cookie(&cookie)
	return nil
}

// Validate Cookie
func ClaimJWT(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	// Get Cookie from headers
	cookie := c.Cookies(config.JWTName)

	// Decoding
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, errors.New("unauthorized")
	}

	// Claim
	claims := token.Claims.(*jwt.StandardClaims)
	return claims, nil
}

func UserFromJWT(c *fiber.Ctx, detailed bool) (interface{}, error) {
	// Claim data from cookie
	claims, err := ClaimJWT(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return nil, err
	}

	// If asked for extra info
	if detailed {
		var user models.User
		database.DB.Where("id = ?", claims.Issuer).First(&user)
		return user, nil
	}

	return claims.Issuer, nil
}
