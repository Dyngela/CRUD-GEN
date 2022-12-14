package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

const baseDirectory = "C:/CRUDGenerator"

func GenerateSpringProject(tables []parser.Table, projectName string) {
	mainDirectory := fmt.Sprintf("%s/%s/src/main/java/com/ne", baseDirectory, strcase.ToLowerCamel(projectName))
	rootDirectory := fmt.Sprintf("%s/%s", baseDirectory, strcase.ToLowerCamel(projectName))
	resourcesDirectory := fmt.Sprintf("%s/%s/src/main/resources", baseDirectory, strcase.ToLowerCamel(projectName))
	exceptionDirectory := fmt.Sprintf("%s/%s/src/main/java/com/ne/exception", baseDirectory, strcase.ToLowerCamel(projectName))

	if err := os.MkdirAll(resourcesDirectory, os.ModePerm); err != nil {
		log.Panic(err)
		return
	}
	if err := os.MkdirAll(exceptionDirectory, os.ModePerm); err != nil {
		log.Panic(err)
		return
	}
	generateJavaMainClass(mainDirectory, projectName)
	generateJavaPomXML(rootDirectory, projectName)
	generateJavaGitignore(rootDirectory)
	//generateJavaApplicationProperties(resourcesDirectory)
	generateJavaException(exceptionDirectory)

	for i := 0; i < len(tables); i++ {
		tableDirectory := fmt.Sprintf("%s/%s/src/main/java/com/ne/%s", baseDirectory, projectName, strcase.ToLowerCamel(tables[i].TableName))
		if err := os.MkdirAll(tableDirectory, os.ModePerm); err != nil {
			log.Panic(err)
			return
		}
		transformDataType(&tables[i])
		generateJavaModel(tables[i], tableDirectory)
		generateJavaDTO(tables[i], tableDirectory)
		generateJavaService(tables[i], tableDirectory)
		generateJavaRepository(tables[i], tableDirectory)
		generateJavaController(tables[i], tableDirectory)
	}

}
