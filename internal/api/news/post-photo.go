package news

import (
	"fmt"
	"strconv"

	"NewsBack/internal/service/news"

	"github.com/gofiber/fiber/v3"
)

func PostPhoto(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Post(url, func(c fiber.Ctx) error {

		file, fileErr := c.FormFile("image")

		if fileErr != nil {
			fmt.Println(fileErr)
			return c.SendStatus(400)
		}

		idInt, errParam := strconv.Atoi(c.Params("id"))

		if errParam != nil {
			fmt.Println(errParam)
			return c.SendStatus(404)
		}

		id := uint(idInt)

		if err := news.SaveImage(c, file, id); err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(201)
	})
}
