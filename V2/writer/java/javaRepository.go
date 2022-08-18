package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

func generateJavaRepository(table parser.Table, path string) {
	var str string
	str = str + generateJavaRepositoryImport(table)
	str = str + fmt.Sprintf(
		`public interface %s extends JpaRepository<%s, %s> {

}`, strcase.ToCamel(table.TableName)+"Repository", strcase.ToCamel(table.TableName), findDataTypeOfPrimitiveTypeAccording(table))
	path = path + "/" + strcase.ToCamel(table.TableName) + "Repository.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaRepositoryImport(table parser.Table) string {
	return fmt.Sprintf(`package com.ne.%s;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
`,
		strcase.ToLowerCamel(table.TableName))
}

func findDataTypeOfPrimitiveTypeAccording(table parser.Table) string {
	for i := 0; i < len(table.Columns); i++ {
		if table.Columns[i].IsPrimaryKey {
			return table.Columns[i].DataType
		}
	}
	log.Panic("No data type found or no primary key")
	return ""
}
