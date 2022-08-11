package java

import (
	"CRUDGEN/V1/src/api/model"
	"CRUDGEN/V1/src/api/repository"
	"fmt"
	"log"
	"os"
)

const baseDirectory = "C:/CRUDGenerator"

func GenerateJavaProject(projectName string) {
	var tables []model.Table
	_, dataErr := repository.FindAllTableAccordingToAProject(&tables, projectName)
	if dataErr != nil {
		return
	}

	// TODO ajouter une relation pour catégoriser les tables.
	resourcesDirectory := fmt.Sprintf("%s/%s/src/main/resources", baseDirectory, projectName)
	modelDirectory := fmt.Sprintf("%s/%s/src/main/java/com/ne/%s/model", baseDirectory, projectName, tables[0].FolderName)
	repositoryDirectory := fmt.Sprintf("%s/%s/src/main/java/com/ne/%s/repository", baseDirectory, projectName, tables[0].FolderName)

	if err := os.MkdirAll(resourcesDirectory, os.ModePerm); err != nil {
		log.Panic(err)
		return
	}
	if err := os.MkdirAll(modelDirectory, os.ModePerm); err != nil {
		log.Panic(err)
		return
	}
	if err := os.MkdirAll(repositoryDirectory, os.ModePerm); err != nil {
		log.Panic(err)
		return
	}

	// TODO A mettre dans des thread différents en fin de projet.
	CreateJavaModel(tables)
	CreateJavaRepositories(tables)
}
