package models

import (
	"gorm.io/gorm"
)

type AuthorRequest struct {
	gorm.Model

	Name string `gorm:"varchar(100);not null" json:"name"`
	Age  uint8  `gorm:"not null" json:"age"`
}

type Author struct {
	AuthorRequest

	Books []Book `gorm:"ForeignKey:AuthorID;association_foreignkey:ID;" json:"books"`
}

type BookAuthorResponse struct {
	AuthorRequest
}

func (BookAuthorResponse) TableName() string {
	return "authors"
}
