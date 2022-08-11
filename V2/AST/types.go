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

type Reference struct {
	ReferenceTable string
	OnDelete       string
	OnUpdate       string
}
