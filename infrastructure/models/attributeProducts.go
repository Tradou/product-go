package models

type AttributeProduct struct {
	ID                uint
	ProductID         uint
	AttributeID       uint
	Value             string    `gorm:"not null;size:64"`
	PriceModification uint      `gorm:"not null"`
	Stock             uint      `gorm:"not null"`
	Product           Product   `gorm:"foreignKey:ProductID"`
	Attribute         Attribute `gorm:"foreignKey:AttributeID"`
}
