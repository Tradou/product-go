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
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	State       bool   `json:"state"`
}
