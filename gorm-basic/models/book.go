package models

import (
	"gorm.io/gorm"
)

type BookRequest struct {
	gorm.Model

	AuthorID int    `gorm:"not null" json:"author_id" binding:"required"`
	Title    string `gorm:"varchar(300);unique;not null" json:"title" binding:"required,max=300"`
	Page     uint32 `gorm:"not null" json:"page" binding:"required,gte=1"`
}

func (BookRequest) TableName() string {
	return "books"
}

type BookResponse struct {
	BookRequest `gorm:"embedded"`

	Author BookAuthorResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"author"`
}

type Book struct {
	BookRequest `gorm:"embedded"`
}
