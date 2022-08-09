package controller

import (
	"CRUDGEN/src/api/services"
	"github.com/gin-gonic/gin"
)

func TableController(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/table/id/:id", services.GetTableById)
		v1.GET("/table/name/:name", services.GetTableByName)
		v1.GET("/tables", services.GetAllTable)
		v1.POST("/table", services.CreateTable)
		v1.POST("/tables", services.CreateTables)
		v1.PUT("/table", services.UpdateTable)
		v1.DELETE("/table/:id", services.DeleteTable)
		v1.DELETE("/table", services.DeleteTableDeletedThreeMonthsAgo)
	}
}
