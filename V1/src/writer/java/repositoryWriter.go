package java

import (
	model2 "CRUDGEN/V1/src/api/model"
	"CRUDGEN/V1/src/writer/writerUtils"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

func CreateJavaRepositories(tables []model2.Table) {
	for i := 0; i < len(tables); i++ {
		f, err := os.Create(
			"C:/CRUDGenerator/myproject/nomDeMicroService/src/main/java/com/ne/nomDeMicroService/repository/" + strcase.ToCamel(tables[i].Name) + "Repository.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		className, _ := createRepo(tables[i])
		_, err2 := f.WriteString(
			className +
				writerUtils.OpeningBracket() +
				JPAMethods(tables[i]) +
				writerUtils.ClosingBracket())

		if err2 != nil {
			log.Fatal(err2)
		}
	}
	fmt.Println("done creating repository")

}

func createRepo(table model2.Table) (string, error) {
	className := strcase.ToCamel(table.Name)
	primaryField := findPrimaryKey(table).(model2.Field)

	repo := fmt.Sprintf(
		`
package com.ne.%s

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface %sRepository extends JpaRepository<%s, %s>`,
		strcase.ToLowerCamel(className),
		strcase.ToCamel(className),
		strcase.ToCamel(className),
		primaryField.Type)
	return repo, nil
}

func JPAMethods(table model2.Table) string {
	primaryField := findPrimaryKey(table).(model2.Field)
	if primaryField.Name == "" {
		return ""
	}

	return fmt.Sprintf(`	Optional<%s> findById(%s id)`, strcase.ToCamel(table.Name), primaryField.Type)
}

func findPrimaryKey(table model2.Table) any {
	for i := 0; i < len(table.Fields); i++ {
		if table.Fields[i].PrimaryKey {
			return table.Fields[i]
		}
	}
	return model2.Field{Name: ""}
}
