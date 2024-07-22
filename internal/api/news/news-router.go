package news

import "github.com/gofiber/fiber/v3"

func NewRoute(url string, App *fiber.App) {
	GetNew(url, "/new/:id", App)
	GetNews(url, "/news", App)
	PostNew(url, "/new", App)
	PatchNew(url, "/new", App)
	DeleteNew(url, "/new/:id", App)
	PostPhoto(url, "/new/imgupload/:id", App)

}
