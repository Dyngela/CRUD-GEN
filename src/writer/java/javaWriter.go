package java

import (
	"CRUDGEN/src/api/model"
	"CRUDGEN/src/api/repository"
	"fmt"
	"log"
	"os"
)

func GenerateJavaProject(projectName string) {
	var tables []model.Table
	_, dataErr := repository.FindAllTable(&tables)
	if dataErr != nil {
		return
	}

	resourcesDirectory := fmt.Sprintf("C:/CRUDGenerator/%s/%s/src/main/resources", projectName, tables[0].FolderName)

	if err := os.MkdirAll("C:/CRUDGenerator/myproject/nomDeMicroService/src/main/resources", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("C:/CRUDGenerator/myproject/nomDeMicroService/src/main/java/com/ne/nomDeMicroService/model", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("C:/CRUDGenerator/myproject/nomDeMicroService/src/main/java/com/ne/nomDeMicroService/repository", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	CreateJavaModel(tables)
	CreateJavaRepositories(tables)
}
