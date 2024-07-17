package tags

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"gorm.io/gorm"
)

func CreateTag(ID []int, NewsID int, Name []string) error {

	var res gorm.DB

	for i := 0; i < len(ID); i++ {

		data, err := GetDataTagsById(ID[i])

		if err != nil {
			return err
		}

		if len(*data) == 0 {

			res := db.DataBase.Create(&models.NewsTags{IDTag: ID[i], IDNews: NewsID})

			if res.Error != nil {
				return res.Error
			}

			res = db.DataBase.Create(&models.Tags{ID: ID[i], Name: Name[i]})

			if res.Error != nil {
				return res.Error
			}
		}
	}

	return res.Error
}
