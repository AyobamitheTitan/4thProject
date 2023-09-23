package controllers

import (
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
)

type ListSerializer struct{
	ListName string `json:"list_name"`
	Songs []SongSerializer
	Books []BookSerializer
	Others []othersSerializer
}

func GetEverything(ctx *gin.Context){
	list_title := ctx.Query("title")
	owner := ctx.Query("owner")

	if list_title == ""{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"A list title must be provided",
		})
		return
	}

	if owner == ""{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"A list owner must be provided",
		})
		return
	}

	var books []models.Book
	var others []models.Other
	var movies []models.Movie
	var songs []models.Song

	go database.DB.DB.Where(&models.Book{Owner: owner}).First(&books,"list_name = ?",list_title)
	go database.DB.DB.Where(&models.Other{Owner: owner}).First(&others,"list_name = ?",list_title)
	database.DB.DB.Where(&models.Song{Owner: owner}).First(&songs,"list_name = ?",list_title)
	database.DB.DB.Where(&models.Movie{Owner: owner}).First(&movies,"list_name = ?",list_title)


	ctx.JSON(http.StatusOK,gin.H{
		"movies":movies,
		"songs":songs,
		"books":books,
		"others":others,
	})
}

