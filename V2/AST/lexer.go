package AST

import (
	"log"
	"os"
	"regexp"
	"strings"
)

var Tables []HandleQuery

func ReadFile() {
	//dat, err := os.ReadFile("C:\\dev\\T&S\\catalogue\\module\\CRUD generator\\CRUD-POC\\V2\\sql\\script.sql")
	//dat, err := os.ReadFile("C:\\dev\\Taff\\T&S\\Catalogue\\CRUDGEN\\V2\\sql\\script.sql")
	dat, err := os.ReadFile("C:\\dev\\Taff\\T&S\\Catalogue\\CRUDGEN\\V2\\sql\\KIS.sql")
	check(err)
	lexFile(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lexFile(sql string) {
	// On split le fichier sql en un tableau d'instructions
	var sqlSplit []string = strings.Split(sql, ";")
	sqlSplit = cleanFile(sqlSplit)

	for i := 0; i < len(sqlSplit); i++ {
		order := strings.ToUpper(sqlSplit[i])

		if strings.Contains(order, "CREATE") {

			index := strings.Index(sqlSplit[i], "(")

			declaration := getCreateTableDeclaration(sqlSplit[i], index)
			setTableName(declaration, i)

			columns := getCreateFieldDeclaration(sqlSplit[i], index)
			setColumns(columns)

		}

		if strings.Contains(order, "DROP") {

		}

		if strings.Contains(order, "ALTER") {

		}
	}
}

func getCreateTableDeclaration(str string, index int) string {
	declaration := str[:index]
	declaration = strings.TrimSpace(declaration)
	return declaration
}

func getCreateFieldDeclaration(str string, index int) []string {
	table := str[index+1:]
	columns := strings.Split(table, ",  ")
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.TrimSpace(columns[i])
	}
	columns[len(columns)-1] = cleanLastParenthesis(columns[len(columns)-1])
	return columns
}

func setTableName(str string, index int) {
	str = cleanDoubleWhiteSpace(str)
	str = strings.TrimSpace(str)
	decomposedDeclaration := strings.Split(str, " ")
	//log.Println(decomposedDeclaration)
	switch strings.ToUpper(decomposedDeclaration[1]) {
	case "TABLE":
		var table HandleQuery
		table.TableName = decomposedDeclaration[len(decomposedDeclaration)-1]
		Tables = append(Tables, table)
	case "AUTRE STATEMENT":
		log.Println(decomposedDeclaration[1])
	default:
		log.Println("Unknown keyword", decomposedDeclaration[1])
	}
}

func setColumns(str []string) {
	var columns []Column
	for i := 0; i < len(str); i++ {
		var primaryKeyFieldName = ""

		var tempColumn = GetColumn()
		cleanColumn := cleanDoubleWhiteSpace(str[i])
		cleanColumn = cleanInParenthesisWhiteSpace(str[i])

		if strings.Contains(cleanColumn, "PRIMARY KEY") {
			primaryKeyFieldName = findConstraintFieldName(cleanColumn)
		}

		if isAColumnWithPrimitiveType(str[i]) {
			tempColumn.DataType = findPrimitiveType(str[i])

			temp := strings.Split(cleanColumn, " ")
			tempColumn.ColumnName = strings.ReplaceAll(temp[0], "`", "")
			if strings.Contains(cleanColumn, "NOT NULL") {
				tempColumn.IsNullable = false
			}
			if strings.Contains(cleanColumn, "UNIQUE") {
				tempColumn.IsUnique = true
			}
			if strings.Contains(cleanColumn, "PRIMARY KEY") {
				tempColumn.IsPrimaryKey = true
			}
			columns = append(columns, tempColumn)
			continue
		}

		if len(primaryKeyFieldName) > 0 {
			fieldName := findConstraintFieldName(cleanColumn)
			log.Println(fieldName)
			tempColumn.IsPrimaryKey = true
		}
		columns = append(columns, tempColumn)

	}
}

//column.IsForeignKey = false
//column.Reference = Reference{}
//column.Length = 0
//column.Precision = 0
//column.DefaultValue = nil

func isAColumnWithPrimitiveType(str string) bool {
	isPrimitive := regexp.MustCompile(
		` int| INT| varchar| VARCHAR| float| FLOAT| datetime| DATETIME| tinyint| TINYINT`)
	return isPrimitive.MatchString(str)
}

func findPrimitiveType(str string) string {
	isPrimitive := regexp.MustCompile(
		` int| INT| varchar| VARCHAR| float| FLOAT| datetime| DATETIME| tinyint| TINYINT`)
	primitiveType := isPrimitive.FindStringSubmatch(str)
	return primitiveType[0]
}

func findConstraintFieldName(str string) string {
	fieldNameFinder := regexp.MustCompile(`\([\s\S]*?\)`)
	fieldName := fieldNameFinder.FindStringSubmatch(str)
	temp := strings.ReplaceAll(fieldName[0], "(`", "")
	return strings.ReplaceAll(temp, "`)", "")
}

func findIndexOfColumnAccordingToItsName() {

}
