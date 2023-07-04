package models

type AttributeProductModel struct{}

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

type GetAttributeProduct struct {
	AttributeProductModel
	ID                uint
	ProductID         uint
	AttributeID       uint
	Value             string `gorm:"not null;size:64"`
	PriceModification uint   `gorm:"not null"`
	Stock             uint   `gorm:"not null"`
}

type StoreAttributeProduct struct {
	AttributeProductModel
	ProductID         uint   `json:"product_id" binding:"required"`
	AttributeID       uint   `json:"attribute_id" binding:"required"`
	Value             string `json:"value" binding:"required"`
	PriceModification uint   `json:"price_modification" binding:"required"`
	Stock             uint   `json:"stock" binding:"required"`
}

type UpdateAttributeProduct struct {
	AttributeProductModel
	Value             string `json:"value"`
	PriceModification uint   `json:"price_modification"`
	Stock             uint   `json:"stock"`
}

type OrderAttributeProduct struct {
	AttributeProductModel
	ID                uint
	ProductID         uint
	AttributeID       uint
	Value             string    `gorm:"not null;size:64"`
	PriceModification uint      `gorm:"not null"`
	Attribute         Attribute `gorm:"foreignKey:AttributeID"`
}

func (AttributeProductModel) TableName() string {
	return "attribute_products"
}
