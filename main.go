package main

import (
	"main/database"
	"main/middlewares"
	"main/routes"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := utils.Init()
	db := database.Connect(config.DbName, config.DbUser, config.DbTcp)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)

	router.Use(middlewares.SetCorsPolicy("http://localhost:5173"))
	router.Use(middlewares.EarlyExitOnPreflightRequests())
	router.Use(middlewares.DBMiddleware(db))
	router.Use(func(context *gin.Context) {
		context.Set("ConfigKey", config)
		context.Next()
	})

	router.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(router)

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}

}
