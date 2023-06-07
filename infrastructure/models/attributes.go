package models

type Attribute struct {
	ID                uint
	Name              string             `gorm:"not null;unique;size:64"`
	AttributeProducts []AttributeProduct `gorm:"foreignKey:AttributeID"`
}
