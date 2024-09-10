package users

import "github.com/gofiber/fiber/v3"

func UserRoute(url string, App *fiber.App) {
	GetUsers(url, "/users", App)
	GetUser(url, "/user/:id", App)
	PostUser(url, "/user", App)
	PatchUser(url, "/user", App)
	DeleteUser(url, "/user/:id", App)

}
