package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajeshj3/jwt-auth/controllers"
	"github.com/rajeshj3/jwt-auth/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/auth/registration/", controllers.Register)
	app.Post("/auth/login/", controllers.Login)

	// following end points requires user to be authenticated
	app.Use(middlewares.AuthMiddleware)

	app.Post("/auth/logout/", controllers.Logout)
	app.Get("/me/", controllers.Me)
}
