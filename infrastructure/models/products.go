package models

import (
	"time"
)

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       uint
	Stock       uint
	State       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
