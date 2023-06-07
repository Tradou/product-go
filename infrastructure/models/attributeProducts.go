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

type StoreAttributeProduct struct {
	ProductID         uint   `json:"product_id" binding:"required"`
	AttributeID       uint   `json:"attribute_id" binding:"required"`
	Value             string `json:"value" binding:"required"`
	PriceModification uint   `json:"price_modification" binding:"required"`
	Stock             uint   `json:"stock" binding:"required"`
}

type UpdateAttributeProduct struct {
	Value             string `json:"value" binding:"required"`
	PriceModification uint   `json:"price_modification" binding:"required"`
	Stock             uint   `json:"stock" binding:"required"`
}

func (StoreAttributeProduct) TableName() string {
	return "attribute_products"
}

func (UpdateAttributeProduct) TableName() string {
	return "attribute_products"
}
