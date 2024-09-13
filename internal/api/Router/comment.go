package Router

import (
	"NewsBack/internal/domain"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
)

type CommentUseCase interface {
	FindAll() ([]domain.Comment, error)
	FindOne(int) (domain.Comment, error)
	Save(domain.Comment) (domain.Comment, error)
	DeleteById(int) (domain.Comment, error)
}

type CommentHandler struct {
	CommentUseCase CommentUseCase
}

func NewCommentRouter(useCase CommentUseCase) *CommentHandler {
	return &CommentHandler{CommentUseCase: useCase}
}

func (uh *CommentHandler) FindAll(c fiber.Ctx) error {
	data, err := uh.CommentUseCase.FindAll()

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *CommentHandler) FindOne(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.CommentUseCase.FindOne(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}

func (uh *CommentHandler) Save(c fiber.Ctx) error {
	body := c.Body()
	var Comment domain.Comment

	err := json.Unmarshal(body, &Comment)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse json")
	}

	if Comment.ID == 0 || Comment.Description == "" || Comment.Name == "" || Comment.UserID == 0 || Comment.NewsID == 0 {
		_ = c.SendStatus(http.StatusBadRequest)
		return c.SendString("data is incorrect or empty")
	}

	data, err := uh.CommentUseCase.Save(Comment)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *CommentHandler) DeleteById(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.CommentUseCase.DeleteById(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}
