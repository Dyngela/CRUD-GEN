package DTO

import "CRUDGEN/src/api/model"

func MapDTOToField(DTO CreateFieldDTO) model.Field {
	var field model.Field
	field.Type = DTO.Type
	field.Name = DTO.Name
	field.Nullable = DTO.Nullable
	field.PrimaryKey = DTO.PrimaryKey
	field.Length = DTO.Length
	field.Precision = DTO.Precision
	field.TableId = DTO.TableId
	return field
}

func MapDTOToUpdateField(DTO UpdateFieldDTO) model.Field {
	var field model.Field
	field.ID = DTO.ID
	field.Type = DTO.Type
	field.Name = DTO.Name
	field.Nullable = DTO.Nullable
	field.PrimaryKey = DTO.PrimaryKey
	field.Length = DTO.Length
	field.Precision = DTO.Precision
	field.TableId = DTO.TableId
	return field
}

func MapFindTableToDTO(field model.Field) UpdateFieldDTO {
	var DTO UpdateFieldDTO
	DTO.ID = field.ID
	DTO.Type = field.Type
	DTO.Name = field.Name
	DTO.Nullable = field.Nullable
	DTO.PrimaryKey = field.PrimaryKey
	DTO.Length = field.Length
	DTO.Precision = field.Precision
	DTO.TableId = field.TableId
	return DTO
}
