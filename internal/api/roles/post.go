package roles

import (
	"NewsBack/internal/service/roles"
	"errors"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PostRole(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Post(url, func(c fiber.Ctx) error {

		err := roles.CreateRole(c.Body())

		if err != nil {
			if errors.Is(err, gorm.DB{}.Error) {
				return c.SendStatus(500)
			} else {
				return c.SendStatus(400)
			}

		} else {
			return c.SendStatus(201)
		}
	})
}
