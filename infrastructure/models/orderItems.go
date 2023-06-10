package models

type OrderItem struct {
	ID                 uint
	OrderID            uint `gorm:"not null"`
	ProductID          uint `gorm:"not null"`
	AttributeProductID uint
	Quantity           uint             `gorm:"not null"`
	Order              Order            `gorm:"foreignKey:OrderID"`
	Product            Product          `gorm:"foreignKey:ProductID"`
	AttributeProduct   AttributeProduct `gorm:"foreignKey:AttributeProductID"`
}
