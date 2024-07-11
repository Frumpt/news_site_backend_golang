package comments

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"encoding/json"
	"fmt"
)

func CreateComment(body []byte) error {
	var JsonData struct {
		ID          *int   `json:"ID"`
		UserID      *int   `json:"UserID"`
		NewsID      *int   `json:"NewsID"`
		Name        string `json:"Name"`
		Description string `json:"Description"`
	}

	if err := json.Unmarshal(body, &JsonData); err != nil {
		return err
	}

	if JsonData.UserID == nil || JsonData.ID == nil || JsonData.NewsID == nil || JsonData.Name == "" || JsonData.Description == "" {
		return fmt.Errorf("name or password or id or user_role_id is empty")
	}

	res := db.DataBase.Create(&models.Comments{ID: *JsonData.ID, UserID: *JsonData.UserID, NewsID: *JsonData.NewsID, Name: JsonData.Name, Description: JsonData.Description})

	return res.Error
}
