package main

import (
	"fmt"
	"os"

	"github.com/arisatriop/go-auto-reload/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Capture database connection
	// config.CreateDBConnection()
	// fmt.Println("Connected to database")

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
