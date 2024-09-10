package roles

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
)

func GetDataRole(id int) (int64, []byte, error, error) {

	var role models.Roles

	res := db.DataBase.Select("ID", "Roles").Find(&role, "id = ?", id)

	data, err := json.Marshal(role)

	return res.RowsAffected, data, err, res.Error
}
