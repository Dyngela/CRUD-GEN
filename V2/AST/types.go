package AST

type HandleQuery struct {
	TableName string
	Columns   []Column
}

type Column struct {
	ColumnName   string
	IsUnique     bool
	IsNullable   bool
	IsPrimaryKey bool
	IsForeignKey bool
	Reference    Reference
	DataType     string
	Length       uint
	Precision    uint
	DefaultValue any
}

func GetColumn() Column {
	var column Column
	column.ColumnName = ""
	column.IsUnique = false
	column.IsNullable = true
	column.IsPrimaryKey = false
	column.IsForeignKey = false
	column.Reference = Reference{}
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
