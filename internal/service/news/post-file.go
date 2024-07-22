package news

import (
	"NewsBack/internal/db"
	"NewsBack/internal/models"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"mime/multipart"
)

func SaveImage(c fiber.Ctx, file *multipart.FileHeader, NewId uint) error {
	destination := fmt.Sprintf("./images/%s", file.Filename)

	if errSave := c.SaveFile(file, destination); errSave != nil {
		return errSave
	}

	if errDB := getNewForImage(file.Filename, NewId); errDB != nil {
		return errDB
	}

	return nil
}

func getNewForImage(NameImage string, IdNew uint) error {
	var newt models.News

	res := db.DataBase.Model(&newt).Where("id = ?", IdNew).Updates(
		&models.News{NameImage: NameImage})

	return res.Error
}
