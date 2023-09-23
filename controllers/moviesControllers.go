package controllers

import (
	"fmt"
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
)

type movieSerializer struct {
	ListName string `json:"list_name"`
	Title    string `json:"title"`
    Remark   string `json:"remark"`
}

func createMovieSerializer(movie models.Movie) movieSerializer{
	return movieSerializer{
		ListName: movie.ListName,
		Title: movie.Title,
		Remark: movie.Remark,
	}
}

func AddMovie(ctx *gin.Context){
	owner := ctx.Query("owner")
	if owner == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An activity owner cannot be empty",
		})	
	}

	var movie models.Movie

	if err:= ctx.BindJSON(&movie); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}	
	movie.Owner = owner
	err := database.DB.DB.Create(&movie).Error

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Unable to add movie to database",
			"details":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"created":true,
		"movie":createMovieSerializer(movie),
	})
}

func GetMovie(ctx *gin.Context)  {
	title := ctx.Query("title")
	owner := ctx.Query("owner")

	if owner == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An activity owner cannot be empty",
		})	
		return
	}

	if title == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An list title cannot be empty",
		})	
		return
	}

	var movies []models.Movie
	var serialized []movieSerializer

	err := database.DB.DB.Where(&models.Movie{Owner: owner}).First(&movies,"list_name = ?",title).Error
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":fmt.Sprintf("The requested list %s does not exist",title),
		})
		return
	}

	for _,m := range movies{
		serialized = append(serialized, createMovieSerializer(m))
	}

	ctx.JSON(http.StatusOK,gin.H{
		"movies":serialized,
	})
}