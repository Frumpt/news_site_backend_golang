package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"fmt"
)

func UpdateNew(body []byte) error {
	var newt models.News

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

	res := db.DataBase.Model(&newt).Where("id = ?", JsonData.ID).Updates(&models.News{Title: JsonData.Title, Description: JsonData.Description, NameImage: JsonData.NameImage})

	return res.Error

}
