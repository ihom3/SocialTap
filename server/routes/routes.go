package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/register-code", controllers.RegisterNewCode)
	app.Post("/api/update-bio", controllers.UpdateBio)
	app.Post("/api/update-email", controllers.UpdateEmail)
	app.Post("/api/update-name", controllers.UpdateName)
	app.Post("/api/update-password", controllers.UpdatePassword)
	app.Post("/api/update-socials", controllers.UpdateSocials)
	app.Get("/api/get-user", controllers.GetUser)
	app.Get("/api/is-logged-in", controllers.IsLoggedIn)
	app.Get("/api/:id", controllers.IDRoute)
	app.Post("/api/update-picture", controllers.UpdateProfilePicture)
	app.Get("/api/profile-picture/:id", controllers.GetProfilePicture)
	//TEMPORARY FOR TESTING
	app.Post("/api/reg-user", controllers.RegisterPersonalUser)
}
