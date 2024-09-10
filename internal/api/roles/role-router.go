package roles

import "github.com/gofiber/fiber/v3"

func RoleRoute(url string, App *fiber.App) {
	GetRole(url, "/Role/:id", App)
	PostRole(url, "/Role", App)
	PatchRole(url, "/Role", App)
	DeleteRole(url, "/Role/:id", App)

}
