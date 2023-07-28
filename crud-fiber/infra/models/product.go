package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	SellerID        int               `json:"seller_id" binding:"required"`
	Name            string            `gorm:"varchar(255);unique;not null" json:"name" binding:"required"`
	ProductQuantity []ProductQuantity `gorm:"ForeignKey:ProductID,association_foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductQuantity struct {
	gorm.Model

	ProductID int    `json:"product_id"`
	Type      string `gorm:"type:enum('IN','OUT')" json:"type" binding:"required"`
	Qty       uint32 `gorm:"not null" json:"qty" binding:"required,gte=1"`
}
