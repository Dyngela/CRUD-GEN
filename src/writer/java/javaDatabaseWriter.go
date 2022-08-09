package java

import (
	"CRUDGEN/src/api/model"
	"fmt"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
)

func CreateJavaModel(tables []model.Table) {
	var str string
	for _, t := range tables {

		log.Println(fieldsToString(t))
		str = fmt.Sprintf(
			`public class %s implements Serializable {
					%s
					}`, cases.Title(language.Und).String(t.Name), fieldsToString(t))
		log.Println(str)
	}

	fe, _ := os.Create("data.txt")
	_, _ = fe.WriteString(str)
}

func fieldsToString(table model.Table) string {
	var fieldsWriter = ""

	for _, f := range table.Fields {
		if f.PrimaryKey == true {
			fieldsWriter = fieldsWriter + fmt.Sprintf("@Id \n@Column(length = %d, unique = true, nullable = false, name = %s)\n", f.Length, strcase.ToSnake(f.Name))
			fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n", strcase.ToCamel(f.Type), strcase.ToLowerCamel(f.Name))
			continue
		}
		fieldsWriter = fieldsWriter + fmt.Sprintf("@Column(length = %d, unique = true, nullable = %v, name = %s)\n", f.Length, f.Nullable, strcase.ToSnake(f.Name))
		fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n", strcase.ToCamel(f.Type), strcase.ToLowerCamel(f.Name))
	}
	return fieldsWriter
}
