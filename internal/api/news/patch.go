package news

import (
	"NewsBack/internal/service/news"
	"github.com/gofiber/fiber/v3"
)

func PatchNew(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Patch(url, func(c fiber.Ctx) error {

		err := news.UpdateNew(c.Body())

		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(200)

	})
}
