package main

import (
	// "log" // for logging the message
	// "net/http" // for core http server in go
	// "github.com/gin-gonic/gin"
	"api/internal/config"
	"api/internal/db"
	"api/internal/logger"
	"api/internal/validator"
	"api/routes"
)

func main() {
	/*
		Starting the server using core
		http server in go
	*/
	// log.Println("Starting server...")

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*
		Making the server using
		gin gonic
	*/
	// router := gin.Default()

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// router.Run(":8080")

	/*
		using router imports
		Handler -> Service -> Repository
	*/

	// Initializing the logger
	logger.InitLogger()
	defer logger.Log.Sync()

	config.LoadConfig()

	validator.InitValidator()

	// Initializing the database
	db.InitDB()

	logger.Log.Info("Server Starting...")

	// Router configuration
	router := routes.SetUpRouter()
	router.Run(":" + config.Appcfg.Port)

}
