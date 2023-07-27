package models

type Product struct {
	Id          int64  `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	ProductName string `gorm:"type:varchar(200);unique;not null" json:"product_name"`
	Decription  string `gorm:"type:text" json:"description"`
}
