package models

type Attribute struct {
	ID                uint
	Name              string             `gorm:"not null;unique;size:64"`
	AttributeProducts []AttributeProduct `gorm:"foreignKey:AttributeID"`
}

type StoreAttribute struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttribute struct {
	Name string `json:"name" binding:"required"`
}
