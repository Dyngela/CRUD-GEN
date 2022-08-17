package AST

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Tables []Table

func ReadFile() {
	//data, err := os.ReadFile("C:\\dev\\Taff\\T&S\\Catalogue\\CRUDGEN\\V2\\sql\\script.sql")
	data, err := os.ReadFile("C:\\dev\\Taff\\T&S\\Catalogue\\CRUDGEN\\V2\\sql\\KIS.sql")
	//data, err := os.ReadFile("C:\\dev\\T&S\\catalogue\\module\\CRUD generator\\CRUD-POC\\V2\\sql\\script.sql")
	//data, err := os.ReadFile("C:\\dev\\T&S\\catalogue\\module\\CRUD generator\\CRUD-POC\\V2\\sql\\KIS.sql")

	check(err)
	mySQLParser(string(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Main parsing method. It parses the SQL and fill an array of Table which will be used to generate code afterward.
*/
func mySQLParser(sql string) {
	// On split le fichier sql en un tableau d'instructions
	var sqlSplit []string = strings.Split(sql, ";")
	sqlSplit = cleanFile(sqlSplit)

	for i := 0; i < len(sqlSplit); i++ {
		order := strings.ToUpper(sqlSplit[i])

		if strings.Contains(order, "CREATE") {

			index := strings.Index(sqlSplit[i], "(")

			declaration := getCreateTableDeclaration(sqlSplit[i], index)
			setTableName(declaration)

			columns := getCreateFieldDeclaration(sqlSplit[i], index)
			setColumns(columns)
		}

		if strings.Contains(order, "DROP") {

		}

		if strings.Contains(order, "ALTER") {

		}
	}
	TablesToString(Tables)
	//TableToString(Tables[0])
}

/* getCreateTableDeclaration
str -> The segment declaration, usually the whole table.
index -> the index of the string where we have to cut the string
Get the table declaration and return it
*/
func getCreateTableDeclaration(str string, index int) string {
	declaration := str[:index]
	declaration = strings.TrimSpace(declaration)
	return declaration
}

/* getCreateFieldDeclaration
str -> The segment declaration, usually the table.
index -> the index of the string where we have to cut the string
Get every field of a specific table and handle EOF
*/
func getCreateFieldDeclaration(str string, index int) []string {
	table := str[index+1:]
	columns := strings.Split(table, ",  ")
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.TrimSpace(columns[i])
	}
	//columns[len(columns)-1] = cleanLastParenthesis(columns[len(columns)-1])
	return columns
}

/* setTableName
str -> the table declaration already separated from its columns
Set the table name
*/
func setTableName(str string) {
	str = cleanDoubleWhiteSpace(str)
	str = strings.TrimSpace(str)
	cleanedDeclaration := strings.ReplaceAll(str, "`", "")
	decomposedDeclaration := strings.Split(cleanedDeclaration, " ")
	switch strings.ToUpper(decomposedDeclaration[1]) {
	case "TABLE":
		// TODO Assert that the last index of table creation is actually the table name
		var table Table
		table.TableName = decomposedDeclaration[len(decomposedDeclaration)-1]
		Tables = append(Tables, table)
	case "AUTRE STATEMENT":
		log.Println(decomposedDeclaration[1])
	default:
		log.Println("Unknown keyword", decomposedDeclaration[1])
	}
}

/* setColumns
str -> All the fields of a table splited to each comma space (",   ")
Fill an array of column with all the desired information.
It's stocked to a table within an array of tables that we use later for generation purposes
*/
func setColumns(str []string) {
	var columns []Column
	for i := 0; i < len(str); i++ {
		var tempColumn = GetColumn()
		// We clean eventual spacing error in sql script
		cleanColumn := cleanDoubleWhiteSpace(str[i])
		cleanColumn = cleanInParenthesisWhiteSpace(str[i])

		if isAColumnWithPrimitiveType(str[i]) {
			tempColumn.DataType = strings.TrimSpace(findPrimitiveType(str[i]))
			length, precision, err := findLengthAndPrecisionOfField(tempColumn.DataType, cleanColumn)
			if err != nil {
				log.Println(err)
				return
			}
			tempColumn.Length = length
			tempColumn.Precision = precision

			// TODO: we get the first word of our string and admit its our column name.
			// We could do it better with a regex for example
			temp := strings.Split(cleanColumn, " ")
			tempColumn.ColumnName = strings.ReplaceAll(temp[0], "`", "")
			if strings.Contains(cleanColumn, " NOT NULL ") {
				tempColumn.IsNullable = false
			}
			if strings.Contains(cleanColumn, " UNIQUE ") {
				tempColumn.IsUnique = true
			}
			if strings.Contains(cleanColumn, " PRIMARY KEY ") {
				tempColumn.IsPrimaryKey = true
			}
			if strings.Contains(cleanColumn, " AUTO_INCREMENT ") {
				tempColumn.IsAutoIncremented = true
			}
			if strings.Contains(cleanColumn, " DEFAULT ") {
				tempColumn.DefaultValue = setDefaultValue(cleanColumn)
			}
			columns = append(columns, tempColumn)
			continue
		}
		Tables[len(Tables)-1].Columns = columns

		// if the occurrence contain PRIMARY KEY but is not with a primitive type, then it's
		// to forward declare the primary key. Then we find the right column according to its name
		// and set its bool IsPrimaryKey to true
		if strings.Contains(cleanColumn, "PRIMARY KEY") {
			primaryKeyFieldName := findConstraintFieldName(cleanColumn)
			setPrimaryKeyToTrueAccordingToColumnName(primaryKeyFieldName)
		}
		if strings.Contains(cleanColumn, "FOREIGN KEY") {
			addForeignKey(str[i])
		}

	}

}

/* findLengthAndPrecisionOfField
valueType -> the type of the field, VARCHAR INT FLOAT etc...
str -> the column itself
Take the column declaration and its type and extract if exists, the length and precision
*/
func findLengthAndPrecisionOfField(valueType string, str string) (uint, uint, error) {
	if len(valueType) <= 0 {
		return 0, 0, nil
	}
	var length uint64
	var precision uint64
	regex := fmt.Sprintf(`%s\([\s\S]*?\)`, valueType)
	fieldNameFinder := regexp.MustCompile(regex)
	fieldName := fieldNameFinder.FindStringSubmatch(str)
	if len(fieldName) == 0 {
		return 0, 0, nil
	}
	fieldLength := strings.ReplaceAll(fieldName[0], valueType, "")
	fieldLength = strings.ReplaceAll(fieldLength, "(", "")
	fieldLength = strings.ReplaceAll(fieldLength, ")", "")
	//
	if strings.Contains(fieldLength, ",") {
		splitString := strings.Split(fieldLength, ",")
		length, _ = strconv.ParseUint(strings.TrimSpace(splitString[0]), 10, 64)
		precision, _ = strconv.ParseUint(strings.TrimSpace(splitString[1]), 10, 64)
		return uint(length), uint(precision), nil
	}
	length, _ = strconv.ParseUint(fieldLength, 10, 64)
	return uint(length), 0, nil
}

/* isAColumnWithPrimitiveType
str -> The column declaration to be treated
Find if a given string contain a primitive type and return true if so.
*/
func isAColumnWithPrimitiveType(str string) bool {
	isPrimitive := regexp.MustCompile(
		` int| INT| varchar| VARCHAR| float| FLOAT| datetime| DATETIME| tinyint| TINYINT`)
	return isPrimitive.MatchString(str)
}

/* findPrimitiveType
str -> The column declaration to be treated
Find if a given string has a primitive type. We usually use it on a column, so we can identify its type
*/
func findPrimitiveType(str string) string {
	isPrimitive := regexp.MustCompile(
		` int| INT| varchar| VARCHAR| float| FLOAT| datetime| DATETIME| tinyint| TINYINT`)
	primitiveType := isPrimitive.FindStringSubmatch(str)
	return primitiveType[0]
}

/* findConstraintFieldName
str -> The column declaration to be treated
Extract a constraint field name with format (`name`) into just name
removing parenthesis and coma
*/
func findConstraintFieldName(str string) string {
	fieldNameFinder := regexp.MustCompile(`\([\s\S]*?\)`)
	fieldName := fieldNameFinder.FindStringSubmatch(str)
	temp := strings.ReplaceAll(fieldName[0], "(`", "")
	return strings.ReplaceAll(temp, "`)", "")
}

/* setPrimaryKeyToTrueAccordingToColumnName
name -> The column's name to be affected by primary key changement
Used to set the primary key of a table if declared forward
*/
func setPrimaryKeyToTrueAccordingToColumnName(name string) {
	for i := 0; i < len(Tables); i++ {
		for x := 0; x < len(Tables[i].Columns); x++ {
			if Tables[i].Columns[x].ColumnName == name {
				Tables[i].Columns[x].IsPrimaryKey = true
			}
		}
	}
}

/* setDefaultValue
str -> the column string which have a default value
Set the default value to a column
*/
func setDefaultValue(str string) string {
	//TODO Handle possible string with space
	defaultValueFinder := regexp.MustCompile(`DEFAULT \s*\S*`)
	fieldName := defaultValueFinder.FindStringSubmatch(str)
	value := strings.ReplaceAll(fieldName[0], "DEFAULT", "")
	return strings.TrimSpace(value)
}

/* addForeignKey
str -> the column string which have foreign key
Add the reference to the table in which the foreign key is originated and call a function to add it to the reverse table
*/
func addForeignKey(str string) {
	defaultValueFinder := regexp.MustCompile(`FOREIGN KEY\s*\(\S*\)`)
	fieldNameRegex := defaultValueFinder.FindStringSubmatch(str)
	if len(fieldNameRegex) == 0 {
		return
	}
	//cleanedField := cleanDoubleWhiteSpace(fieldNameRegex[0])
	splitField := strings.Split(fieldNameRegex[0], " ")
	fieldName := strings.ReplaceAll(splitField[2], "(", "")
	fieldName = strings.ReplaceAll(fieldName, ")", "")
	fieldName = strings.ReplaceAll(fieldName, "`", "")
	// TODO Assert that the first occurrence is the field of the table, and the second the field in the foreign table
	var reference Reference
	// We declare the field created with the same name as the foreign key like so.
	for x := 0; x < len(Tables[len(Tables)-1].Columns); x++ {
		if Tables[len(Tables)-1].Columns[x].ColumnName == fieldName {
			reference.ReferenceTable, reference.FieldName = findForeignTableNameAndField(str)
			reference.MappingType = "OneToMany"
			reference.OnUpdate = findOnUpdateType(str)
			reference.OnDelete = findOnDeleteType(str)
			Tables[len(Tables)-1].Columns[x].IsForeignKey = true
			Tables[len(Tables)-1].Columns[x].Reference = append(Tables[len(Tables)-1].Columns[x].Reference, reference)
		}
	}
	addReferenceToForeignTable(str)

}

/* addReferenceToForeignTable
str -> the column string which have foreign key with reference.
Add a reference to the reverse table from foreign key declaration.
*/
func addReferenceToForeignTable(str string) {
	var foreignReference Reference
	findReference := regexp.MustCompile(`REFERENCES \s*\S*\s*\(.*\)`)
	reference := findReference.FindStringSubmatch(str)

	cleanedField := cleanDoubleWhiteSpace(reference[0])
	splitField := strings.Split(cleanedField, " ")
	tableName := strings.ReplaceAll(splitField[1], "`", "")

	columnName := strings.ReplaceAll(splitField[2], "(", "")
	columnName = strings.ReplaceAll(columnName, ")", "")
	columnName = strings.ReplaceAll(columnName, "`", "")

	referenceTable := Tables[len(Tables)-1].TableName
	foreignReference.ReferenceTable = referenceTable
	foreignReference.FieldName = columnName
	foreignReference.MappingType = "ManyToOne"

	for i := 0; i < len(Tables); i++ {
		if Tables[i].TableName == tableName {
			for x := 0; x < len(Tables[i].Columns); x++ {
				if Tables[i].Columns[x].ColumnName == columnName {
					Tables[i].Columns[x].Reference = append(Tables[i].Columns[x].Reference, foreignReference)
				}
			}
		}
	}
}

/* findOnUpdateType
str -> the column string which have foreign key with reference.
Find if there's an ON UPDATE order for a particular foreign key and store it
*/
func findOnUpdateType(str string) string {
	defaultValueFinder := regexp.MustCompile(`ON\s*UPDATE\s*[A-Z]*`)
	fieldName := defaultValueFinder.FindStringSubmatch(str)
	if len(fieldName) == 0 {
		return ""
	}
	cleanedFieldName := cleanDoubleWhiteSpace(fieldName[0])
	splitField := strings.Split(cleanedFieldName, " ")
	return splitField[2]
}

/* findOnDeleteType
str -> the column string which have foreign key with reference.
Find if there's an ON DELETE order for a particular foreign key and store it
*/
func findOnDeleteType(str string) string {
	defaultValueFinder := regexp.MustCompile(`ON\s*DELETE\s*[A-Z]*`)
	fieldName := defaultValueFinder.FindStringSubmatch(str)
	if len(fieldName) == 0 {
		return ""
	}
	cleanedFieldName := cleanDoubleWhiteSpace(fieldName[0])
	splitField := strings.Split(cleanedFieldName, " ")
	return splitField[2]
}

/* findForeignTableNameAndField
find the occurrence of the table and its column corresponding to the reverse table of a particular foreign key
*/
func findForeignTableNameAndField(str string) (string, string) {
	findReference := regexp.MustCompile(`REFERENCES \s*\S*\s*\(.*\)`)
	reference := findReference.FindStringSubmatch(str)
	cleanedField := cleanDoubleWhiteSpace(reference[0])
	splitField := strings.Split(cleanedField, " ")
	tableName := strings.ReplaceAll(splitField[1], "`", "")

	findKey := regexp.MustCompile(`FOREIGN KEY\s*\(\S*\)`)
	key := findKey.FindStringSubmatch(str)
	cleanedField = cleanDoubleWhiteSpace(key[0])
	splitField = strings.Split(key[0], " ")
	fieldName := strings.ReplaceAll(splitField[2], "(", "")
	fieldName = strings.ReplaceAll(fieldName, ")", "")
	fieldName = strings.ReplaceAll(fieldName, "`", "")

	return tableName, fieldName
}
