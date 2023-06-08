package models

type AttributeModel struct{}

type Attribute struct {
	ID                uint
	Name              string             `gorm:"not null;unique;size:64"`
	AttributeProducts []AttributeProduct `gorm:"foreignKey:AttributeID"`
}

type GetAttribute struct {
	AttributeModel
	ID   uint
	Name string `gorm:"not null;unique;size:64"`
}

type StoreAttribute struct {
	AttributeModel
	Name string `json:"name" binding:"required"`
}

type UpdateAttribute struct {
	AttributeModel
	Name string `json:"name" binding:"required"`
}

func (AttributeModel) TableName() string {
	return "attributes"
}
