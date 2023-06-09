package models

import "time"

type Order struct {
	ID          uint
	OrderNumber string      `gorm:"not null;unique"`
	TotalAmount uint        `gorm:"not null"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}