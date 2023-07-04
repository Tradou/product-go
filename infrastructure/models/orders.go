package models

import "time"

type Order struct {
	ID         uint
	Amount     uint        `gorm:"not null"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type GetOrder struct {
	ID         uint
	Amount     uint           `gorm:"not null"`
	OrderItems []GetOrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time      `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

func (GetOrder) TableName() string {
	return "orders"
}
