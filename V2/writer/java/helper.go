package java

import (
	"CRUDGEN/V2/parser"
)

func transformDataType(table *parser.Table) {
	for i := 0; i < len(table.Columns); i++ {
		switch table.Columns[i].DataType {
		case "INT":
			table.Columns[i].DataType = "Long"
		case "int":
			table.Columns[i].DataType = "Long"
		case "FLOAT":
			table.Columns[i].DataType = "Float"
		case "float":
			table.Columns[i].DataType = "Float"
		case "VARCHAR":
			table.Columns[i].DataType = "String"
		case "varchar":
			table.Columns[i].DataType = "String"
		case "TINYINT":
			table.Columns[i].DataType = "Boolean"
		case "tinyint":
			table.Columns[i].DataType = "Boolean"
		case "DATETIME":
			table.Columns[i].DataType = "LocalDateTime"
		case "datetime":
			table.Columns[i].DataType = "LocalDateTime"
		}
	}
}
