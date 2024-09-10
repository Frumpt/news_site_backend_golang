package Router

import (
	"NewsBack/internal/domain"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
	"strconv"
)

type UserUseCase interface {
	FindAll() ([]domain.User, error)
	FindOne(int) (domain.User, error)
	Save(domain.User) (domain.User, error)
	DeleteById(int) (domain.User, error)
}

type Response struct {
	ID         int     `gorm:"column:id;primary_key;unique" json:"ID"`
	UserRoleID int     `gorm:"column:user_role_id;NOT NULL" json:"UserRoleID"`
	Name       string  `gorm:"column:name;NOT NULL" json:"Name"`
	Password   *string `gorm:"column:password;NOT NULL" json:"Password,omitempty"`
}

type Handler struct {
	userUseCase UserUseCase
}

func NewUserRouter(useCase UserUseCase) *Handler {
	return &Handler{userUseCase: useCase}
}

func (uh *Handler) FindAll(c fiber.Ctx) error {
	data, err := uh.userUseCase.FindAll()

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *Handler) FindOne(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.userUseCase.FindOne(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}

func (uh *Handler) Save(c fiber.Ctx) error {
	body := c.Body()
	var user domain.User

	err := json.Unmarshal(body, &user)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse json")
	}

	if user.ID == 0 || user.Password == nil || user.Name == "" {
		_ = c.SendStatus(http.StatusBadRequest)
		return c.SendString("data is incorrect or empty")
	}

	data, err := uh.userUseCase.Save(user)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(data)
}

func (uh *Handler) DeleteById(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.SendString("can not parse id")
	}

	data, err := uh.userUseCase.DeleteById(id)

	if err != nil {
		_ = c.SendString(err.Error())
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(data)
}
