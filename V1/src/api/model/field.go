package model

import "gorm.io/gorm"

type Field struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);" json:"name"`
	Type       string `gorm:"type:varchar(255);" json:"type"`
	Length     int    `gorm:"type:integer;" json:"length"`
	PrimaryKey bool   `gorm:"type:boolean;" json:"primaryKey"`
	Nullable   bool   `gorm:"type:boolean;" json:"nullable"`
	Precision  int    `gorm:"type:integer;" json:"precision"`
	Table      Table
	TableId    uint `gorm:"column:tableId;not null;" json:"tableId"`
}
