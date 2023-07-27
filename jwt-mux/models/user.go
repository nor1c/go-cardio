package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Fullname string `gorm:"varchar(200)" json:"fullname"`
	Username string `gorm:"varchar(16),unique" json:"username"`
	Password string `gorm:varchar(200) json:"password"`
}
