package AST

import "log"

type Table struct {
	TableName string
	Columns   []Column
}

func TableToString(table Table) {
	log.Printf(`tableName: %s`, table.TableName)
	for i := 0; i < len(table.Columns); i++ {
		log.Printf(`ColumnName: %s 
			IsUnique    	  : %v
			IsNullable  	  : %v
			IsPrimaryKey	  : %v
			IsAutoIncremented : %v
			IsForeignKey	  : %v
			Reference   	  : %v
			DataType    	  : %s
			Length      	  : %d
			Precision   	  : %d
			DefaultValue	  : %v`,
			table.Columns[i].ColumnName, table.Columns[i].IsUnique, table.Columns[i].IsNullable,
			table.Columns[i].IsPrimaryKey, table.Columns[i].IsAutoIncremented,
			table.Columns[i].IsForeignKey, table.Columns[i].Reference,
			table.Columns[i].DataType, table.Columns[i].Length, table.Columns[i].Precision, table.Columns[i].DefaultValue)
	}
	log.Println("----------------------------")
}

func TablesToString(tables []Table) {
	for i := 0; i < len(tables); i++ {
		TableToString(tables[i])
	}
}

type Column struct {
	ColumnName        string
	IsAutoIncremented bool
	IsUnique          bool
	IsNullable        bool
	IsPrimaryKey      bool
	IsForeignKey      bool
	Reference         []Reference
	DataType          string
	Length            uint
	Precision         uint
	DefaultValue      any
}

func GetColumn() Column {
	var column Column
	column.ColumnName = ""
	column.IsAutoIncremented = false
	column.IsUnique = false
	column.IsNullable = true
	column.IsPrimaryKey = false
	column.IsForeignKey = false
	column.Reference = []Reference{}
	column.DataType = ""
	column.Length = 0
	column.Precision = 0
	column.DefaultValue = nil
	return column
}

type Reference struct {
	ReferenceTable string
	OnDelete       string
	OnUpdate       string
}
