package utils

import (
	model2 "CRUDGEN/V1/src/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=gernika54 dbname=CRUD port=5432"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres")
	}
}
func SyncDatabase() {
	err := Db.AutoMigrate(&model2.Table{}, &model2.Field{})
	CheckForError(err, "Problem automigrating data")
}
