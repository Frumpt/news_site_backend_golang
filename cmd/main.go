package main

import (
	"NewsBack/internal/api/comments"
	"NewsBack/internal/api/news"
	"NewsBack/internal/api/roles"
	"NewsBack/internal/api/users"
	"NewsBack/internal/db"
	serviceUser "NewsBack/internal/service/users"
	"github.com/gofiber/fiber/v3"
	"log"
)

var App *fiber.App

var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func main() {
	App = fiber.New()

	db.Connect(configDB)

	users.UserRoute("/api", App)

	App.Use(serviceUser.UserIndifity)
	roles.RoleRoute("/api", App)
	news.NewRoute("/api", App)
	comments.CommentRoute("/api", App)

	log.Fatal(App.Listen(":3000"))
}
