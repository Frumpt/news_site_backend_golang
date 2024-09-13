package domain

type User struct {
	ID         int     `gorm:"column:id;primary_key;unique" json:"ID"`
	UserRoleID int     `gorm:"column:user_role_id;NOT NULL" json:"UserRoleID"`
	Name       string  `gorm:"column:name;NOT NULL" json:"Name"`
	Password   *string `gorm:"column:password;NOT NULL" json:"Password,omitempty"`
}
