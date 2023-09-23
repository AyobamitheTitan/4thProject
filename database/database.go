package database

import (
	"log"
	"os"
	"v1/listee/config"
	"v1/listee/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInstance struct{
	DB *gorm.DB
}

var DB DBInstance;

func ConnnectDB(){
	godotenv.Load()
	db,err := gorm.Open(postgres.Open(config.NewDsn()),&gorm.Config{});
	if err != nil{
		log.Fatal("Unable to connect to database : \n",err.Error())
		os.Exit(1)
	}

	log.Println("Connected to the database")
	
	log.Println("Running migrations...")
	db.AutoMigrate(&models.Book{},&models.Movie{},&models.Other{},&models.Song{},&models.User{})
	DB = DBInstance{DB: db}
}