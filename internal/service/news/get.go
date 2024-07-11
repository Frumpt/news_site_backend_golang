package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
)

func GetDataNews() ([]byte, error) {
	var news []models.News

	db.DataBase.Select("ID", "UserID", "Title", "Description", "NameImage").Find(&news)

	data, err := json.Marshal(news)
	if err != nil {
		return nil, err
	}

	return data, err
}

func GetDataNew(id int) (int64, []byte, error, error) {
	var newt models.News

	res := db.DataBase.Select("Id", "Name").Find(&newt, "id = ?", id)

	data, err := json.Marshal(newt)

	return res.RowsAffected, data, err, res.Error
}
