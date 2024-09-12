package main

import (
	"NewsBack/internal/api"
	"NewsBack/internal/api/Router"
	db2 "NewsBack/internal/db"
	"NewsBack/internal/repository"
	"NewsBack/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"log"
)

var App *fiber.App

var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func main() {

	db, err := db2.Connect(configDB)
	if err != nil {
		panic(err)
	}

	dbPGX, err := db2.ConnectPGX(configDB)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(dbPGX)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := Router.NewUserRouter(userUseCase)

	newsRepository := repository.NewNewsRepository(db)
	newsUseCase := usecase.NewNewsUseCase(newsRepository)
	newsHandler := Router.NewNewsRouter(newsUseCase)

	commentRepository := repository.NewCommentRepository(db)
	commentUseCase := usecase.NewCommentUseCase(commentRepository)
	commentHandler := Router.NewCommentRouter(commentUseCase)

	tagRepository := repository.NewTagRepository(db)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	tagHandler := Router.NewTagRouter(tagUseCase)
	serverHTTP := api.NewServerHTTP(userHandler, newsHandler, commentHandler, tagHandler)
	log.Fatal(serverHTTP.Listen(":3000"))

}
