package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"fmt"
)

func CreateNew(body []byte) error {
	var JsonData struct {
		ID          *int   `json:"ID"`
		UserID      *int   `json:"UserID"`
		Title       string `json:"Title"`
		Description string `json:"Description"`
		NameImage   string `json:"NameImage"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return err
	}

	if JsonData.UserID == nil || JsonData.ID == nil || JsonData.Title == "" || JsonData.Description == "" || JsonData.NameImage == "" {
		return fmt.Errorf("name or password or id or user_role_id is empty")
	}

	res := db.DataBase.Create(&models.News{ID: *JsonData.ID, UserID: *JsonData.UserID,
		Title: JsonData.Title, Description: JsonData.Description, NameImage: JsonData.NameImage})

	return res.Error
}
