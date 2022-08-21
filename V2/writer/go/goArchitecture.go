package _go

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

const baseDirectory = "C:/CRUDGenerator"

func GenerateGinProject(tables []parser.Table, projectName string) {
	mainDirectory := fmt.Sprintf("%s/%s/src", baseDirectory, strcase.ToLowerCamel(projectName))
	rootDirectory := fmt.Sprintf("%s/%s", baseDirectory, strcase.ToLowerCamel(projectName))

	generateGoMainClass(mainDirectory, projectName)
	generateGoGitignore(rootDirectory)
	generateGoDatabaseLinker(mainDirectory)
	generateGoUtilsDirectory(mainDirectory)

	for i := 0; i < len(tables); i++ {
		tableDirectory := fmt.Sprintf("%s/%s/src/%s",
			baseDirectory, projectName, strcase.ToLowerCamel(tables[i].TableName))
		if err := os.MkdirAll(tableDirectory, os.ModePerm); err != nil {
			log.Panic(err)
			return
		}
		transformDataType(&tables[i])
		generateGoModel(tables[i], tableDirectory)
		generateGoDTO(tables[i], tableDirectory)
		generateGoService(tables[i], tableDirectory)
		generateGoRepository(tables[i], tableDirectory)
		generateGoController(tables[i], tableDirectory)
	}
}