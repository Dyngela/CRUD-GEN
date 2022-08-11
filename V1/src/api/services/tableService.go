package services

import (
	"CRUDGEN/V1/src/api/DTO"
	"CRUDGEN/V1/src/api/model"
	"CRUDGEN/V1/src/api/repository"
	"CRUDGEN/V1/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTable(c *gin.Context) {
	var tableDTO DTO.CreateTableDTO
	var table model.Table
	if err := c.ShouldBindJSON(&tableDTO); err != nil {
		utils.CheckForError(err, "error binding json for table creation")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data sent"})
		return
	}
	table = DTO.MapCreateTableDTOToTable(tableDTO)
	status, err := repository.CreateTable(&table)
	if err != nil {
		c.JSON(status, gin.H{"status": status, "error": err})
		return
	}
	c.JSON(status, gin.H{"message": "Table successfully created"})
}

func CreateTables(c *gin.Context) {
	var tableDTO []DTO.CreateTableDTO
	var table model.Table
	var status int
	var err error

	if err = c.ShouldBindJSON(&tableDTO); err != nil {
		utils.CheckForError(err, "error binding json for table creation")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data sent"})
		return
	}

	for i := 0; i < len(tableDTO); i++ {
		table = DTO.MapCreateTableDTOToTable(tableDTO[i])
		status, err = repository.CreateTable(&table)
		if err != nil {
			c.JSON(status, gin.H{"status": status, "error": err})
			return
		}
	}

	c.JSON(status, gin.H{"message": "Table successfully created"})
}

func UpdateTable(c *gin.Context) {
	var tableDTO DTO.UpdateTableDTO
	var table model.Table
	if err := c.ShouldBindJSON(&tableDTO); err != nil {
		utils.CheckForError(err, "error binding json for table creation")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong data sent"})
		return
	}
	table = DTO.MapUpdateTableDTOToTable(tableDTO)
	status, err := repository.UpdateTable(&table)
	if err != nil {
		c.JSON(status, gin.H{"status": status, "error": err})
		return
	}
	c.JSON(status, gin.H{"message": "Table successfully updated"})
}

func GetTableById(c *gin.Context) {
	var table model.Table
	status, err := repository.FindTableById(&table, c.Param("id"))

	if err != nil {
		c.JSON(status, gin.H{
			"status": http.StatusNoContent,
			"error":  "We couldn't find your table",
		})
		return
	}
	c.JSON(status, DTO.MapTableToFindTableDTO(table))
}

func GetTableByName(c *gin.Context) {
	var table model.Table
	status, err := repository.FindTableByName(&table, c.Param("name"))

	if err != nil {
		c.JSON(status, gin.H{
			"status": http.StatusNoContent,
			"error":  "We couldn't find your table",
		})
		return
	}
	c.JSON(status, DTO.MapTableToFindTableDTO(table))
}

func GetAllTable(c *gin.Context) {
	var tables []model.Table
	status, err := repository.FindAllTable(&tables)

	if err != nil {
		c.JSON(status, gin.H{
			"status": http.StatusNoContent,
			"error":  "We couldn't find all tables",
		})
		return
	}
	c.JSON(status, DTO.MapTablesToFindTableDTO(tables))
}

func DeleteTable(c *gin.Context) {
	status, err := repository.DeleteTable(c.Param("id"))
	if err != nil {
		c.JSON(status, gin.H{
			"status": http.StatusNoContent,
			"error":  "We couldn't delete your user",
		})
		return
	}
	c.JSON(status, gin.H{"message": "Successfully deleted your table"})
}

func DeleteTableDeletedThreeMonthsAgo(c *gin.Context) {
	status, err := repository.DeleteDefinitelyTable()
	if err != nil {
		c.JSON(status, gin.H{
			"status": http.StatusNoContent,
			"error":  "We couldn't delete users with 3 months or more delete time",
		})
		return
	}
	c.JSON(status, gin.H{"message": "Successfully definitely deleted all tables deleted 3 months ago"})
}
