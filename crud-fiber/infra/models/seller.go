package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model

	Name     string    `gorm:"varchar(255);" json:"name" validate:"required"`
	Address  string    `gorm:"type:text;not null;" json:"address" validate:"required"`
	IsActive string    `gorm:"type:enum('0','1');default:0;" json:"is_active"`
	Products []Product `gorm:"ForeignKey:SellerID;association_foreignKey:ID;" json:"products"`
}
