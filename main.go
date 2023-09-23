package main

import (
	"v1/listee/database"
	"v1/listee/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	database.ConnnectDB()
	router := gin.Default()
	
	routes.SetupRoutes(router)
	router.Run("localhost:9798")
}