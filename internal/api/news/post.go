package news

import (
	"NewsBack/internal/service/news"
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PostNew(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Post(url, func(c fiber.Ctx) error {
		err := news.CreateNew(c.Body())

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
