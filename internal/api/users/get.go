package users

import (
	"NewsBack/internal/service/users"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func GetUsers(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Get(url, func(c fiber.Ctx) error {
		data, err := users.GetDataUsers()
		if err != nil {
			return c.SendStatus(500)
		}
		return c.Send(data)
	})
}

func GetUser(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Get(url, func(c fiber.Ctx) error {
		id, errParam := strconv.Atoi(c.Params("id"))
		if errParam != nil {
			return c.SendStatus(404)
		}

		finderId := uint(id)

		rows, data, err, errRes := users.GetDataUser(finderId)

		if rows == 0 {
			return c.SendStatus(404)
		}

		if err != nil && errRes != nil {
			return c.SendStatus(500)
		}

		return c.Send(data)
	})
}
