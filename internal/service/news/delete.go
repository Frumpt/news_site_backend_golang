package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/tags"
	"fmt"
	"os"
)

func DeleteNew(id int) (int64, error) {
	var newt models.News

	resGet := db.DataBase.Select("NameImage").Find(&newt, "id = ?", id)

	if resGet.Error != nil {
		return 0, resGet.Error
	}

	var NameImage string = newt.NameImage

	res := db.DataBase.Where("id = ?", id).Delete(&newt)

	if res.Error != nil {
		return 0, res.Error
	}

	CountRows, err := tags.DeleteTag(id)

	if CountRows == 0 || err != nil {
		return 0, err
	}

	destination := fmt.Sprintf("./images/%s", NameImage)

	if errRemove := os.Remove(destination); errRemove != nil {
		fmt.Println(errRemove)
	}

	return res.RowsAffected, res.Error
}
