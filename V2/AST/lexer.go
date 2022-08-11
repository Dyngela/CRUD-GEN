package AST

import (
	"log"
	"os"
	"strings"
)

func ReadFile() {
	dat, err := os.ReadFile("C:\\dev\\T&S\\catalogue\\module\\CRUD generator\\CRUD-POC\\V2\\sql\\script.sql")
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
			_ = getCreateTableDeclaration(sqlSplit[i], index)
			_ = getCreateFieldDeclaration(sqlSplit[i], index)

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
	columns := strings.Split(table, ",    ")
	for i := 0; i < len(columns); i++ {
		columns[i] = strings.TrimSpace(columns[i])
	}
	columns[len(columns)-1] = cleanLastParenthesis(columns[len(columns)-1])
	log.Println(columns[len(columns)-1])

	return columns
}
