package roles

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"errors"
)

func UpdateRole(body []byte) (int64, error) {

	var role models.Roles

	var JsonData struct {
		ID   *int   `json:"id"`
		Name string `json:"name"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return 0, err
	}

	if JsonData.Name == "" || JsonData.ID == nil {
		return 0, errors.New("name or password is empty")
	}

	res := db.DataBase.Model(&role).Where("id = ?", JsonData.ID).Updates(&models.Roles{Roles: JsonData.Name})

	return res.RowsAffected, res.Error

}
