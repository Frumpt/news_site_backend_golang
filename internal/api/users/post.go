package users

import (
	"NewsBack/internal/service/users"
	"github.com/gofiber/fiber/v3"
)

func PostUser(url string, addUrl string, App *fiber.App) {
	url = url + addUrl

	App.Post(url, func(c fiber.Ctx) error {

		err := users.CreateUser(c.Body())

		if err != nil {
			switch err.Error() {
			case "ERROR: duplicate key value violates unique constraint \"uni_users_id\" (SQLSTATE 23505)":
				return c.SendStatus(400)
			case "name or password or id or user_role_id is empty":
				return c.SendStatus(400)

			case "role not found":
				return c.SendStatus(400)
			}
		}
		return c.SendStatus(201)
	})
}
