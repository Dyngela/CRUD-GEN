package repository

import (
	"CRUDGEN/src/api/model"
	"CRUDGEN/src/utils"
)

func CreateTable(table *model.Table) (int, error) {
	err := utils.Db.Create(&table).Error
	return utils.CheckForQueryError(err, "Error Create user")
}

func UpdateTable(table *model.Table) (int, error) {
	err := utils.Db.Save(&table).Error
	if err != nil {
		return utils.CheckForQueryError(err, "Error with bcrypt")
	}
	return utils.CheckForQueryError(err, "Error Update table")
}

func FindTableById(table *model.Table, id string) (int, error) {
	err := utils.Db.Preload("Fields").First(&table, id).Error
	return utils.CheckForQueryError(err, "Error FindTableById")
}

func FindTableByName(table *model.Table, name string) (int, error) {
	err := utils.Db.Preload("Fields").Where("name = ?", name).First(&table).Error
	return utils.CheckForQueryError(err, "Error FindTableByName")
}

func FindAllTable(tables *[]model.Table) (int, error) {
	err := utils.Db.Preload("Fields").Find(&tables).Error
	return utils.CheckForQueryError(err, "Error FindUserById")
}

func DeleteTable(id string) (int, error) {
	err := utils.Db.Delete(&model.Table{}, id).Error
	return utils.CheckForQueryError(err, "Error Delete user")
}

func DeleteDefinitelyTable() (int, error) {
	err := utils.Db.Unscoped().Where("deleted_at < now() - interval '3 months'").Delete(&model.Table{}).Error
	return utils.CheckForQueryError(err, "Error Deleting permanently user")
}
