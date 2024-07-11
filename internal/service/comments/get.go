package comments

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
)

func GetDataComments() ([]byte, error) {
	var comments []models.Comments

	db.DataBase.Select("ID", "UserID", "Title", "Description", "NameImage").Find(&comments)

	data, err := json.Marshal(comments)
	if err != nil {
		return nil, err
	}

	return data, err
}

func GetDataComment(id int) (int64, []byte, error, error) {
	var comment models.Comments

	res := db.DataBase.Select("Id", "Name").Find(&comment, "id = ?", id)

	data, err := json.Marshal(comment)

	return res.RowsAffected, data, err, res.Error
}
