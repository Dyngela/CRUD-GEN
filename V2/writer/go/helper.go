package _go

import (
	"CRUDGEN/V2/parser"
)

func transformDataType(table *parser.Table) {
	for i := 0; i < len(table.Columns); i++ {
		switch table.Columns[i].DataType {
		case "INT":
			table.Columns[i].DataType = "int"
		case "int":
			table.Columns[i].DataType = "int"
		case "FLOAT":
			table.Columns[i].DataType = "float64"
		case "float":
			table.Columns[i].DataType = "float64"
		case "VARCHAR":
			table.Columns[i].DataType = "string"
		case "varchar":
			table.Columns[i].DataType = "string"
		case "TINYINT":
			table.Columns[i].DataType = "bool"
		case "tinyint":
			table.Columns[i].DataType = "bool"
		case "DATETIME":
			table.Columns[i].DataType = "time.Time"
		case "datetime":
			table.Columns[i].DataType = "time.Time"
		}
	}
}
