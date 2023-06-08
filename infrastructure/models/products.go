package models

import (
	_ "gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint
	Name        string      `gorm:"not null;size:64"`
	Description string      `gorm:"not null"`
	Price       uint        `gorm:"not null"`
	Stock       uint        `gorm:"not null"`
	State       bool        `gorm:"not null"`
	CreatedAt   time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	Attributes  []Attribute `gorm:"many2many:attribute_products;"`
}

type StoreProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Stock       uint   `json:"stock" binding:"required"`
	State       *bool  `json:"state" binding:"required"`
}

type UpdateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	State       *bool  `json:"state"`
}

func (StoreProduct) TableName() string {
	return "products"
}

func (UpdateProduct) TableName() string {
	return "products"
}
