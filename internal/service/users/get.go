package users

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
)

func GetDataUsers() ([]byte, error) {
	var users []models.Users

	db.DataBase.Select("Id", "Name", "user_role_id").Find(&users)

	data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return data, err
}

func GetDataUser(id uint) (int64, []byte, error, error) {
	var user models.Users

	res := db.DataBase.Select("Id", "Name").Find(&user, "id = ?", id)

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	return res.RowsAffected, data, err, res.Error
}
