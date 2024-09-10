package comments

import (
	"NewsBack/internal/service/comments"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func DeleteComment(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Delete(url, func(c fiber.Ctx) error {
		id, errParam := strconv.Atoi(c.Params("id"))
		if errParam != nil {
			return c.SendStatus(400)
		}

		data, err := comments.DeleteComment(id)

		if err != nil {
			return c.SendStatus(500)
		}

		if data == 0 {
			return c.SendStatus(404)
		}

		return c.SendStatus(200)
	})
}
