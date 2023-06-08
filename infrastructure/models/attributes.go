package models

type Attribute struct {
	ID                uint
	Name              string             `gorm:"not null;unique;size:64"`
	AttributeProducts []AttributeProduct `gorm:"foreignKey:AttributeID"`
}

type GetAttribute struct {
	ID   uint
	Name string `gorm:"not null;unique;size:64"`
}

type StoreAttribute struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttribute struct {
	Name string `json:"name" binding:"required"`
}

func (GetAttribute) TableName() string {
	return "attributes"
}

func (StoreAttribute) TableName() string {
	return "attributes"
}

func (UpdateAttribute) TableName() string {
	return "attributes"
}
