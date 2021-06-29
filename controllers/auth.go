package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rajeshj3/jwt-auth/database"
	"github.com/rajeshj3/jwt-auth/models"
	"github.com/rajeshj3/jwt-auth/responses"
	"github.com/rajeshj3/jwt-auth/utils"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// copy data into an interface
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// create user object
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
	}

	// clean data
	user.Prepare()

	// validate data
	err := user.Validate("register")
	if err != nil {
		return responses.ErrorResponse(c, err)
	}

	// save user
	savedUser, err := user.SaveUser(database.DB)
	if err != nil {
		return responses.ErrorResponse(c, err)
	}

	// SetCookie [HttpOnly]
	utils.CreateJWTCookie(c, user.ID)

	return responses.SuccessResponse(c, savedUser)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// copy data into an interface
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// create user instance
	var user models.User

	// search with email
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		return responses.ErrorResponse(c, errors.New("invalid email"))
	}

	// match passwords
	err := models.VerifyPassword(user.Password, data["password"])
	if err != nil {
		return responses.ErrorResponse(c, errors.New("invalid password"))
	}

	// SetCookie [HttpOnly]
	utils.CreateJWTCookie(c, user.ID)

	return responses.SuccessResponse(c, user)
}

func Logout(c *fiber.Ctx) error {
	// SetCookie BLANK [HttpOnly]
	utils.CreateBlankJWTCookie(c)
	return c.SendStatus(fiber.StatusOK)
}

func Me(c *fiber.Ctx) error {

	// get user_id from cookie
	// get user from database
	user, _ := utils.UserFromJWT(c, true)

	return responses.SuccessResponse(c, user)
}
