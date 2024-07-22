package comments

import "github.com/gofiber/fiber/v3"

func CommentRoute(url string, App *fiber.App) {
	GetComments(url, "/comments", App)
	GetComment(url, "/comment/:id", App)
	PostComment(url, "/comment", App)
	PatchComment(url, "/comment", App)
	DeleteComment(url, "/comment/:id", App)
}
