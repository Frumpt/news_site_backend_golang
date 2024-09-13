package domain

type News struct {
	ID          int    `gorm:"column:id;primary_key;unique" json:"ID"`
	UserID      int    `gorm:"column:user_id;NOT NULL" json:"UserID"`
	Title       string `gorm:"column:title;NOT NULL" json:"Title"`
	Description string `gorm:"column:description;NOT NULL" json:"Description"`
	NameImage   string `gorm:"column:name_image;NOT NULL" json:"NameImage"`
}
