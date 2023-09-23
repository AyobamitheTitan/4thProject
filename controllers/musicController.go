 package controllers

import (
	"fmt"
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
)

type SongSerializer struct{
	ListName string `json:"listName"`
    Title    string	`json:"title"`
    Artist   string	`json:"artist"`
    Remark   string	`json:"remark"`
}

func createSongSerializer(song models.Song) SongSerializer{
	return SongSerializer{
		Title: song.Title,
		ListName: song.ListName,
		Artist: song.Artist,
		Remark: song.Remark,
	}
}

func AddSong(ctx *gin.Context){
	owner := ctx.Query("owner")
	if owner == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An activity owner cannot be empty",
		})	
	}

	var song models.Song

	if err := ctx.BindJSON(&song); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	song.Owner = owner
	err := database.DB.DB.Create(&song).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Unable to add song to database",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"added":true,
		"song": createSongSerializer(song),
	})
}

func GetSong(ctx * gin.Context){
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

	var songs []models.Song
	var serialized []SongSerializer

	err := database.DB.DB.Where(&models.Song{Owner: owner}).First(&songs,"list_name = ?",title).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":fmt.Sprintf("The List %s could not be found",title),
		})
		return
	}

	for _,s := range songs{
		serialized = append(serialized, createSongSerializer(s))
	}

	ctx.JSON(http.StatusOK,gin.H{
		"songs": serialized,
	})
}