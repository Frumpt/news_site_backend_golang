package comments

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func DeleteComment(id int) (int64, error) {
	var comment models.Comments

	res := db.DataBase.Where("id = ?", id).Delete(&comment)

	return res.RowsAffected, res.Error
}
