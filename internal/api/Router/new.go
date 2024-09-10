package Router

import (
	"NewsBack/internal/domain"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
)

type NewsUseCase interface {
	FindAll() ([]domain.News, error)
	FindOne(int) (domain.News, error)
	Save(domain.News) (domain.News, error)
	DeleteById(int) (domain.News, error)
}

type NewsHandler struct {
	NewsUseCase NewsUseCase
}

func NewNewsRouter(useCase NewsUseCase) *NewsHandler {
	return &NewsHandler{NewsUseCase: useCase}
}

func (uh *NewsHandler) FindAll(c fiber.Ctx) error {
	data, err := uh.NewsUseCase.FindAll()

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *NewsHandler) FindOne(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.NewsUseCase.FindOne(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}

func (uh *NewsHandler) Save(c fiber.Ctx) error {
	body := c.Body()
	var News domain.News

	err := json.Unmarshal(body, &News)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse json")
	}

	if News.ID == 0 || News.UserID == 0 || News.Description == "" || News.Title == "" {
		_ = c.SendStatus(http.StatusBadRequest)
		return c.SendString("data is incorrect or empty")
	}

	data, err := uh.NewsUseCase.Save(News)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *NewsHandler) DeleteById(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.NewsUseCase.DeleteById(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}
