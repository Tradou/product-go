package models

import (
	"time"
)

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	State       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StoreProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	State       *bool  `json:"state" binding:"required"`
}

func (StoreProduct) TableName() string {
	return "products"
}
