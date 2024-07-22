package users

import (
	"NewsBack/internal/service/users"
	"github.com/gofiber/fiber/v3"
)

func SingUp(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Get(url, func(c fiber.Ctx) error {

		token, err := users.SingUp(c.Body())

		if err != nil {
			return c.SendStatus(500)
		}

		return c.Send(token)
	})
}
