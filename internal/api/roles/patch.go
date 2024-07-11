package roles

import (
	"NewsBack/internal/service/roles"
	"github.com/gofiber/fiber/v3"
)

func PatchRole(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Patch(url, func(c fiber.Ctx) error {

		res, err := roles.UpdateRole(c.Body())

		if res == 0 {
			return c.SendStatus(404)
		}
		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(200)

	})
}
