package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(255);" json:"name"`
	FolderName string  `gorm:"type:varchar(255);" json:"folderName"`
	Fields     []Field `gorm:"constraint:OnDelete:CASCADE;" json:"fields"`
}
