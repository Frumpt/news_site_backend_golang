package tags

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
)

func GetDataTagsByNewsId(NewsID int) ([]string, error) {
	var tags []models.NewsTags

	db.DataBase.Select("id_tag", "id_news").Find(&tags, "id_news = ?", NewsID)

	tagsNames := make([]string, 0, len(tags))

	for _, tag := range tags {
		var oneTag models.Tags
		db.DataBase.Select("Name").Find(&oneTag, "id = ?", tag.IDTag)
		tagsNames = append(tagsNames, oneTag.Name)
	}

	return tagsNames, nil
}

func GetDataTagsById(IDTag int) (*[]string, error) {
	var tags []models.NewsTags

	res := db.DataBase.Select("id_tag", "id_news").Find(&tags, "id_tag = ?", IDTag)

	if res.Error != nil {
		return nil, res.Error
	}

	tagsNames := make([]string, 0, len(tags))

	for _, tag := range tags {
		var oneTag models.Tags
		db.DataBase.Select("Name").Find(&oneTag, "id = ?", tag.IDTag)

		if res.Error != nil {
			return nil, res.Error
		}

		tagsNames = append(tagsNames, oneTag.Name)
	}

	return &tagsNames, nil
}
