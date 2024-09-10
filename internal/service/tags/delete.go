package tags

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"fmt"
	"gorm.io/gorm/clause"
)

func DeleteTag(NewId int) (int64, error) {

	var tag models.Tags

	var newsTag []models.NewsTags

	res := db.DataBase.Clauses(clause.Returning{}).Where("id_news = ?", NewId).Delete(&newsTag)

	if res.Error != nil {
		return 0, res.Error
	}

	fmt.Println(len(newsTag))

	for i := 0; i < len(newsTag); i++ {

		tagId := newsTag[i].IDTag

		res = db.DataBase.Find(&newsTag, "id_tag = ?", tagId)

		if res.Error != nil {
			return 0, res.Error
		}

		if res.RowsAffected == 0 {
			res = db.DataBase.Where("ID = ?", tagId).Delete(&tag)
			if res.Error != nil {
				return 0, res.Error
			}

		}

	}

	return res.RowsAffected, res.Error

}
