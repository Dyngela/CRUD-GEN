package main

import (
	"CRUDGEN/V1/src/api/controller"
	utils2 "CRUDGEN/V1/src/utils"
	"CRUDGEN/V1/src/writer/java"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func init() {
	utils2.ConnectToDatabase()
	utils2.SyncDatabase()
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
}

func main() {

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	utils2.CheckForError(err, "Problem with proxies in main")

	controller.TableController(router)

	java.GenerateJavaProject("POCRUD")

	err = router.Run(port)
	utils2.CheckForError(err, "Fatal error with router")
}
