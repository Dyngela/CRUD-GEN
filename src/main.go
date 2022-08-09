package main

import (
	"CRUDGEN/src/api/controller"
	"CRUDGEN/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func init() {
	utils.ConnectToDatabase()
	utils.SyncDatabase()
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
}

func main() {
	fmt.Println("Server starting")

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	utils.CheckForError(err, "Problem with proxies in main")

	controller.TableController(router)

	err = router.Run(port)
	utils.CheckForError(err, "Fatal error with router")
}
