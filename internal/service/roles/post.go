package roles

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"errors"
)

func CreateRole(body []byte) error {
	var JsonData struct {
		ID   *int   `json:"id"`
		Name string `json:"name"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return err
	}

	if JsonData.Name == "" || JsonData.ID == nil {
		return errors.New("name or password is empty")
	}

	res := db.DataBase.Create(&models.Roles{Roles: JsonData.Name, ID: *JsonData.ID})

	return res.Error
}
