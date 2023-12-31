package main

import (
	"fmt"
	"os"
	"v1/listee/database"
	"v1/listee/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main(){
	godotenv.Load()
	database.ConnnectDB()
	router := gin.Default()
	
	routes.SetupRoutes(router)
	router.Run(fmt.Sprintf(":%s",os.Getenv("PORT")))
}