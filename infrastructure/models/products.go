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
