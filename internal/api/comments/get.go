package comments

import (
	"NewsBack/internal/service/comments"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func GetComments(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Get(url, func(c fiber.Ctx) error {
		data, err := comments.GetDataComments()
		if err != nil {
			return c.SendStatus(500)
		}
		return c.Send(data)
	})
}

func GetComment(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Get(url, func(c fiber.Ctx) error {
		id, errParam := strconv.Atoi(c.Params("id"))
		if errParam != nil {
			return c.SendStatus(404)
		}

		rows, data, err, errRes := comments.GetDataComment(id)

		if rows == 0 {
			return c.SendStatus(404)
		}

		if err != nil && errRes != nil {
			return c.SendStatus(500)
		}

		return c.Send(data)
	})
}
