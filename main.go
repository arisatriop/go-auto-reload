package main

import (
	"fmt"
	"os"

	"github.com/arisatriop/goilerplate-mono/config"
	"github.com/arisatriop/goilerplate-mono/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Capture database connection
	config.CreateDBConnection()

	// Run server
	router := gin.Default()
	routes.InitRoute(router)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("=========================================================================================")
	fmt.Println("=================================== SVC READY TO SERVE ==================================")
	fmt.Println("=========================================================================================")
	fmt.Println("")
	fmt.Println("")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router.Run(":" + port)
}
