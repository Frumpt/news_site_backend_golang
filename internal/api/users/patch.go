package users

import (
	"NewsBack/internal/service/users"
	"github.com/gofiber/fiber/v3"
)

func PatchUser(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Patch(url, func(c fiber.Ctx) error {

		res, err := users.UpdateUser(c.Body())

		if res == 0 {
			return c.SendStatus(404)
		}
		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(200)

	})
}
