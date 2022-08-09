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
	_, dataErr := repository.FindAllTableAccordingToAProject(&tables, projectName)
	if dataErr != nil {
		return
	}

	// TODO ajouter une relation pour définir le projet et catégoriser les tables.
	resourcesDirectory := fmt.Sprintf("C:/CRUDGenerator/%s/%s/src/main/resources", projectName, tables[0].FolderName)
	modelDirectory := fmt.Sprintf("C:/CRUDGenerator/%s/%s/src/main/java/com/ne/%s/model", projectName, tables[0].FolderName, tables[0].FolderName)
	repositoryDirectory := fmt.Sprintf("C:/CRUDGenerator/%s/%s/src/main/java/com/ne/%s/repository", projectName, tables[0].FolderName, tables[0].FolderName)

	if err := os.MkdirAll(resourcesDirectory, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(modelDirectory, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(repositoryDirectory, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// TODO A mettre dans des thread différents en fin de projet.
	CreateJavaModel(tables)
	CreateJavaRepositories(tables)
}
