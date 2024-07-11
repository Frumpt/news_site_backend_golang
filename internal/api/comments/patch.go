package comments

import (
	"NewsBack/internal/service/comments"
	"github.com/gofiber/fiber/v3"
)

func PatchComment(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Patch(url, func(c fiber.Ctx) error {

		err := comments.UpdateComment(c.Body())

		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(200)

	})
}
