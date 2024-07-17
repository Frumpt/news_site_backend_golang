package comments

import (
	"NewsBack/internal/service/comments"
	"errors"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func PostComment(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Post(url, func(c fiber.Ctx) error {

		err := comments.CreateComment(c.Body())

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
