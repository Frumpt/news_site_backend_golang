package users

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func DeleteUser(id uint) (int64, error) {
	var user models.Users

	res := db.DataBase.Where("id = ?", id).Delete(&user)

	return res.RowsAffected, res.Error
}
