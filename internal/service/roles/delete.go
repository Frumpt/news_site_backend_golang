package roles

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

var role models.Roles

func DeleteRole(id int) (int64, error) {

	res := db.DataBase.Where("id = ?", id).Delete(&role)

	return res.RowsAffected, res.Error
}
