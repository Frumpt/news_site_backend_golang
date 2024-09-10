package domain

type Comment struct {
	ID          int    `gorm:"column:id;primary_key" json:"ID"`
	UserID      int    `gorm:"column:user_id;NOT NULL" json:"UserID"`
	NewsID      int    `gorm:"column:news_id;NOT NULL" json:"NewsID"`
	Name        string `gorm:"column:name;NOT NULL" json:"Name"`
	Description string `gorm:"column:description;NOT NULL" json:"Description"`
}
