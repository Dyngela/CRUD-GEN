package repository

import (
	"CRUDGEN/V1/src/api/model"
	utils2 "CRUDGEN/V1/src/utils"
)

func CreateTable(table *model.Table) (int, error) {
	err := utils2.Db.Create(&table).Error
	return utils2.CheckForQueryError(err, "Error Create user")
}

func UpdateTable(table *model.Table) (int, error) {
	err := utils2.Db.Save(&table).Error
	if err != nil {
		return utils2.CheckForQueryError(err, "Error with bcrypt")
	}
	return utils2.CheckForQueryError(err, "Error Update table")
}

func FindTableById(table *model.Table, id string) (int, error) {
	err := utils2.Db.Preload("Fields").First(&table, id).Error
	return utils2.CheckForQueryError(err, "Error FindTableById")
}

func FindTableByName(table *model.Table, name string) (int, error) {
	err := utils2.Db.Preload("Fields").Where("name = ?", name).First(&table).Error
	return utils2.CheckForQueryError(err, "Error FindTableByName")
}

func FindAllTable(tables *[]model.Table) (int, error) {
	err := utils2.Db.Preload("Fields").Find(&tables).Error
	return utils2.CheckForQueryError(err, "Error FindUserById")
}

func FindAllTableAccordingToAProject(tables *[]model.Table, projectName string) (int, error) {
	err := utils2.Db.Preload("Fields").Where("project_name = ?", projectName).Find(&tables).Error
	return utils2.CheckForQueryError(err, "Error FindUserById")
}

func DeleteTable(id string) (int, error) {
	err := utils2.Db.Delete(&model.Table{}, id).Error
	return utils2.CheckForQueryError(err, "Error Delete user")
}

func DeleteDefinitelyTable() (int, error) {
	err := utils2.Db.Unscoped().Where("deleted_at < now() - interval '3 months'").Delete(&model.Table{}).Error
	return utils2.CheckForQueryError(err, "Error Deleting permanently user")
}
