package domain

type Tag struct {
	ID   int    `gorm:"column:id;primary_key;unique" json:"ID"`
	Name string `gorm:"column:name;NOT NULL" json:"Name"`
}
