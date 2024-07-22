package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"NewsBack/internal/service/tags"
	"encoding/json"
	"fmt"
)

func CreateNew(body []byte) error {
	type JsonDataType struct {
		ID          *int      `json:"ID"`
		UserID      *int      `json:"UserID"`
		Title       string    `json:"Title"`
		Description string    `json:"Description"`
		NameImage   string    `json:"NameImage"`
		IDTag       *[]int    `json:"IDTag"`
		NameTag     *[]string `json:"NameTag"`
	}

	var JsonGet JsonDataType

	if err := json.Unmarshal(body, &JsonGet); err != nil {
		return err
	}

	var JsonData JsonDataType = JsonGet

	if JsonData.UserID == nil || JsonData.ID == nil || JsonData.Title == "" || JsonData.Description == "" || JsonData.NameImage == "" {
		return fmt.Errorf("name or password or id or user_role_id is empty")
	}

	res := db.DataBase.Create(&models.News{ID: *JsonData.ID, UserID: *JsonData.UserID,
		Title: JsonData.Title, Description: JsonData.Description, NameImage: JsonData.NameImage})

	if res.Error != nil {
		return res.Error
	}

	if JsonData.IDTag != nil && JsonData.NameTag != nil || JsonData.ID != nil {
		err := tags.CreateTag(*JsonData.IDTag, *JsonData.ID, *JsonData.NameTag)
		if err != nil {
			return err
		}
	}

	return res.Error
}
