package models

type Users struct {
	ID         int     `gorm:"column:id;primary_key;unique" json:"ID"`
	UserRoleID int     `gorm:"column:user_role_id;NOT NULL" json:"UserRoleID"`
	Name       string  `gorm:"column:name;NOT NULL" json:"Name"`
	Password   *string `gorm:"column:password;NOT NULL" json:"Password,omitempty"`
}

type Comments struct {
	ID          int    `gorm:"column:id;primary_key" json:"ID"`
	UserID      int    `gorm:"column:user_id;NOT NULL" json:"UserID"`
	NewsID      int    `gorm:"column:news_id;NOT NULL" json:"NewsID"`
	Name        string `gorm:"column:name;NOT NULL" json:"Name"`
	Description string `gorm:"column:description;NOT NULL" json:"Description"`
}

type Roles struct {
	ID    int    `gorm:"column:id;primary_key" json:"ID"`
	Roles string `gorm:"column:roles;NOT NULL" json:"Roles"`
}

type News struct {
	ID          int    `gorm:"column:id;primary_key;unique" json:"ID"`
	UserID      int    `gorm:"column:user_id;NOT NULL" json:"UserID"`
	Title       string `gorm:"column:title;NOT NULL" json:"Title"`
	Description string `gorm:"column:description;NOT NULL" json:"Description"`
	NameImage   string `gorm:"column:name_image;NOT NULL" json:"NameImage"`
}

type NewsTags struct {
	IDNews int `gorm:"column:id_news;NOT NULL" json:"IDNews"`
	IDTag  int `gorm:"column:id_tag;NOT NULL" json:"IDTag"`
}

type Tags struct {
	ID   int    `gorm:"column:id;primary_key;unique" json:"ID"`
	Name string `gorm:"column:name;NOT NULL" json:"Name"`
}
