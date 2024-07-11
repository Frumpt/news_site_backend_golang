package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func DeleteNew(id int) (int64, error) {
	var new models.News

	res := db.DataBase.Where("id = ?", id).Delete(&new)

	return res.RowsAffected, res.Error
}
