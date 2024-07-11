package users

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/roles"
	"encoding/json"
	"fmt"
)

func CreateUser(body []byte) error {
	var JsonData struct {
		ID         *int   `json:"id"`
		UserRoleID *int   `json:"user_role_id"`
		Name       string `json:"name"`
		Password   string `json:"password"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return err
	}

	if JsonData.Name == "" || JsonData.Password == "" || JsonData.ID == nil || JsonData.UserRoleID == nil {
		return fmt.Errorf("name or password or id or user_role_id is empty")
	}

	countRows, data, errJSON, err := roles.GetDataRole(*JsonData.UserRoleID)

	if countRows == 0 || errJSON != nil || err != nil || data == nil {
		return fmt.Errorf("role not found")
	}
	res := db.DataBase.Create(&models.Users{Name: JsonData.Name, Password: &JsonData.Password, ID: *JsonData.ID,
		UserRoleID: *JsonData.UserRoleID})

	return res.Error
}
