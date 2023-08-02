package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	SellerID        int               `json:"seller_id" binding:"required"`
	Name            string            `gorm:"varchar(255);unique;not null" json:"name" binding:"required"`
	ProductQuantity []ProductQuantity `gorm:"ForeignKey:ProductID;association_foreignkey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"quantities"`
	AvailableStock  AvailableStock    `gorm:"ForeignKey:ProductID;association_foreignkey:ID;" json:"stock"`
}

type ProductQuantity struct {
	gorm.Model

	ProductID int    `json:"product_id"`
	Type      string `gorm:"type:enum('IN','OUT')" json:"type" binding:"required"`
	Qty       uint32 `gorm:"not null" json:"qty" binding:"required,gte=1"`
}

type AvailableStock struct {
	ProductID    int    `json:"product_id"`
	CurrentStock uint32 `json:"current_stock"`
}

func (AvailableStock) TableName() string {
	return "available_stock"
}
