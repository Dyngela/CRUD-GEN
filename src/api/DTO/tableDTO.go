package DTO

import (
	"CRUDGEN/src/api/model"
)

type CreateTableDTO struct {
	Name       string           `binding:"required" json:"name"`
	FolderName string           `json:"folderName"`
	Fields     []CreateFieldDTO `binding:"required" json:"fields"`
}
type CreateFieldDTO struct {
	Name       string `binding:"required" json:"name"`
	Type       string `binding:"required" json:"type"`
	Length     int    `binding:"required" json:"length"`
	PrimaryKey bool   `binding:"required" json:"primaryKey"`
	Nullable   bool   `binding:"required" json:"nullable"`
	Precision  int    `binding:"required" json:"precision"`
	TableId    uint   `binding:"required" json:"tableId"`
}

func MapCreateTableDTOToTable(DTO CreateTableDTO) model.Table {
	var table model.Table
	var fields []model.Field
	table.Name = DTO.Name
	table.FolderName = DTO.FolderName
	for i := 0; i < len(DTO.Fields); i++ {
		fields = append(fields, MapDTOToField(DTO.Fields[i]))
	}
	table.Fields = fields
	return table
}

type UpdateTableDTO struct {
	ID         uint             `binding:"required" json:"id"`
	Name       string           `binding:"required" json:"name"`
	FolderName string           `binding:"required" json:"folderName"`
	Fields     []UpdateFieldDTO `binding:"required" json:"fields"`
}
type UpdateFieldDTO struct {
	ID         uint   `binding:"required" json:"id"`
	Name       string `binding:"required" json:"name"`
	Type       string `binding:"required" json:"type"`
	Length     int    `binding:"required" json:"length"`
	PrimaryKey bool   `binding:"required" json:"primaryKey"`
	Nullable   bool   `binding:"required" json:"nullable"`
	Precision  int    `binding:"required" json:"precision"`
	TableId    uint   `binding:"required" json:"tableId"`
}

func MapUpdateTableDTOToTable(DTO UpdateTableDTO) model.Table {
	var table model.Table
	var fields []model.Field
	table.ID = DTO.ID
	table.Name = DTO.Name
	table.FolderName = DTO.FolderName
	for i := 0; i < len(DTO.Fields); i++ {
		fields = append(fields, MapDTOToUpdateField(DTO.Fields[i]))
	}
	table.Fields = fields
	return table
}

type FindTableDTO struct {
	ID         uint             `binding:"required" json:"id"`
	Name       string           `binding:"required" json:"name"`
	FolderName string           `binding:"required" json:"folderName"`
	Fields     []UpdateFieldDTO `binding:"required" json:"fields"`
}
type FindFieldDTO struct {
	ID         uint   `binding:"required" json:"id"`
	Name       string `binding:"required" json:"name"`
	Type       string `binding:"required" json:"type"`
	Length     int    `binding:"required" json:"length"`
	PrimaryKey bool   `binding:"required" json:"primaryKey"`
	Nullable   bool   `binding:"required" json:"nullable"`
	Precision  int    `binding:"required" json:"precision"`
	TableId    uint   `binding:"required" json:"tableId"`
}

func MapTableToFindTableDTO(table model.Table) FindTableDTO {
	var DTO FindTableDTO
	var fieldsDTO []UpdateFieldDTO
	DTO.ID = table.ID
	DTO.Name = table.Name
	DTO.FolderName = table.FolderName

	for i := 0; i < len(table.Fields); i++ {
		fieldsDTO = append(fieldsDTO, MapFindTableToDTO(table.Fields[i]))
	}
	DTO.Fields = fieldsDTO
	return DTO
}

func MapTablesToFindTableDTO(table []model.Table) []FindTableDTO {
	var DTO []FindTableDTO
	for i := 0; i < len(table); i++ {
		DTO = append(DTO, MapTableToFindTableDTO(table[i]))
	}
	return DTO
}
