package users

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"errors"
)

func UpdateUser(body []byte) (int64, error) {

	var user models.Users

	var JsonData struct {
		Id       uint   `json:"Id"`
		Name     string `json:"Name"`
		Password string `json:"Password"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return 0, err
	}

	if JsonData.Name == "" || JsonData.Password == "" {
		return 0, errors.New("name or password is empty")
	}

	res := db.DataBase.Model(&user).Where("id = ?", JsonData.Id).Updates(models.Users{Name: JsonData.Name,
		Password: &JsonData.Password})

	return res.RowsAffected, res.Error

}
