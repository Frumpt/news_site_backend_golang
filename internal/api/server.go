package api

import (
	"NewsBack/internal/api/Router"
	"github.com/gofiber/fiber/v3"
)

type serverHTTP struct {
	app *fiber.App
}

func NewServerHTTP(userHandler *Router.Handler, newsHandler *Router.NewsHandler, commentHandler *Router.CommentHandler, tagHandler *Router.TagHandler) *fiber.App {
	app := fiber.New()
	app.Get("/users", userHandler.FindAll)
	app.Get("/user/:id", userHandler.FindOne)
	app.Post("/user", userHandler.Save)
	app.Delete("/user/:id", userHandler.DeleteById)

	app.Get("/news", newsHandler.FindAll)
	app.Get("/new/:id", newsHandler.FindOne)
	app.Post("/new", newsHandler.Save)
	app.Delete("/new/:id", newsHandler.DeleteById)

	app.Get("/comments", commentHandler.FindAll)
	app.Get("/comment/:id", commentHandler.FindOne)
	app.Post("/comment", commentHandler.Save)
	app.Delete("/comment/:id", commentHandler.DeleteById)

	app.Get("/tags", tagHandler.FindAll)
	app.Get("/tag/:id", tagHandler.FindOne)
	app.Post("/tag", tagHandler.Save)
	app.Delete("/tag/:id", tagHandler.DeleteById)
	return app
}
