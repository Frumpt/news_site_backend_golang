package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func DeleteNew(id int) (int64, error) {
	var newt models.News

	res := db.DataBase.Where("id = ?", id).Delete(&newt)

	return res.RowsAffected, res.Error
}
