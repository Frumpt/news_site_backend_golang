package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/tags"
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

	res := db.DataBase.Select("ID", "UserID", "Title", "Description", "NameImage").Find(&newt, "id = ?", id)

	dataTags, errTags := tags.GetDataTagsByNewsId(id)

	if errTags != nil {
		return 0, nil, errTags, res.Error
	}

	dataWithTags := struct {
		DataTags []string
		Data     models.News
	}{dataTags, newt}

	data, err := json.Marshal(dataWithTags)

	return res.RowsAffected, data, err, res.Error
}
