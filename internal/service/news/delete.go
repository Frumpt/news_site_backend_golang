package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/tags"
)

func DeleteNew(id int) (int64, error) {
	var newt models.News

	res := db.DataBase.Where("id = ?", id).Delete(&newt)

	if res.Error != nil {
		return 0, res.Error
	}

	CountRows, err := tags.DeleteTag(id)

	if CountRows == 0 || err != nil {
		return 0, err
	}

	return res.RowsAffected, res.Error
}
