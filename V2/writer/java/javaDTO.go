package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

func generateJavaDTO(table parser.Table, path string) {
	var str string
	fields, tables := generateJavaFieldsDTO(table)
	str = str + generateJavaDTOImport(table, tables)
	str = str + fmt.Sprintf(
		`public class %s {
	%s
}`, strcase.ToCamel(table.TableName), fields)

	path = path + "/" + table.TableName + "DTO.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaDTOImport(table parser.Table, tables []string) string {
	var formattedTablesImport string

	for i := 0; i < len(tables); i++ {
		formattedTablesImport = formattedTablesImport +
			fmt.Sprintf("import com.ne.%s.%s;\n",
				strcase.ToLowerCamel(tables[i]), strcase.ToCamel(tables[i]))
	}

	return fmt.Sprintf(`package com.ne.%s;

%s

import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
`,
		strcase.ToLowerCamel(table.TableName), formattedTablesImport)
}

func generateJavaFieldsDTO(table parser.Table) (string, []string) {
	var fieldsWriter string
	var relationWriter string
	var classToImport []string

	for i := 0; i < len(table.Columns); i++ {
		fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s;\n\t",
			table.Columns[i].DataType, strcase.ToLowerCamel(table.Columns[i].ColumnName))

		if len(table.Columns[i].Reference) > 0 {
			for ref := 0; ref < len(table.Columns[i].Reference); ref++ {
				classToImport = append(classToImport, strcase.ToCamel(table.Columns[i].Reference[ref].ReferenceTable)+"DTO")
				if table.Columns[i].Reference[ref].MappingType == "OneToMany" {
					log.Println("aze")

					classType := strcase.ToCamel(table.Columns[i].Reference[ref].ReferenceTable) + "DTO"
					relationWriter = relationWriter + fmt.Sprintf("private List<%s> %s;\n\t",
						classType, strcase.ToLowerCamel(table.Columns[i].Reference[ref].ReferenceTable))
				}
				if table.Columns[i].Reference[ref].MappingType == "ManyToOne" {
					classType := strcase.ToCamel(table.Columns[i].Reference[ref].ReferenceTable) + "DTO"
					relationWriter = relationWriter + fmt.Sprintf("private %s %s;\n\t",
						classType, strcase.ToLowerCamel(table.Columns[i].Reference[ref].ReferenceTable))
				}
			}
		}
	}

	fieldsWriter = fieldsWriter + "\n\t" + relationWriter

	return fieldsWriter, classToImport
}
