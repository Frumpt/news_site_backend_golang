package Router

import (
	"NewsBack/internal/domain"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
)

type TagUseCase interface {
	FindAll() ([]domain.Tag, error)
	FindOne(int) (domain.Tag, error)
	Save(domain.Tag) (domain.Tag, error)
	DeleteById(int) (domain.Tag, error)
}

type TagHandler struct {
	TagUseCase TagUseCase
}

func NewTagRouter(useCase TagUseCase) *TagHandler {
	return &TagHandler{TagUseCase: useCase}
}

func (uh *TagHandler) FindAll(c fiber.Ctx) error {
	data, err := uh.TagUseCase.FindAll()

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *TagHandler) FindOne(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.TagUseCase.FindOne(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}

func (uh *TagHandler) Save(c fiber.Ctx) error {
	body := c.Body()
	var Tag domain.Tag

	err := json.Unmarshal(body, &Tag)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse json")
	}

	if Tag.ID == 0 || Tag.Name == "" {
		_ = c.SendStatus(http.StatusBadRequest)
		return c.SendString("data is incorrect or empty")
	}

	data, err := uh.TagUseCase.Save(Tag)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *TagHandler) DeleteById(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.TagUseCase.DeleteById(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}
