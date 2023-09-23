package controllers

import (
	"fmt"
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
)

type othersSerializer struct {
	ListName string `json:"list_name"`
	Title    string `json:"title"`
    Remark   string `json:"remark"`
}

func createOthersSerializer(others models.Other) othersSerializer{
	return othersSerializer{
		ListName: others.ListName,
		Title: others.Title,
		Remark: others.Remark,
	}
}

func AddOther(ctx *gin.Context){
	owner := ctx.Query("owner")
	if owner == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An activity owner cannot be empty",
		})	
	}

	var others models.Other

	if err:= ctx.BindJSON(&others); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}	

	others.Owner = owner
	err := database.DB.DB.Create(&others).Error

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Unable to add to database",
			"details":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"created":true,
		"others":createOthersSerializer(others),
	})
}

func GetOther(ctx *gin.Context)  {
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

	var others []models.Other
	var serialized []othersSerializer

	err := database.DB.DB.Where(&models.Other{Owner: owner}).First(&others,"list_name = ?",title).Error
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":fmt.Sprintf("The requested list %s does not exist",title),
		})
		return
	}

	for _,o := range others{
		serialized = append(serialized, createOthersSerializer(o))
	}

	ctx.JSON(http.StatusOK,gin.H{
		"others":serialized,
	})
}