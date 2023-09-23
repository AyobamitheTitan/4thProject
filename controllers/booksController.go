package controllers

import (
	"fmt"
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
)

type BookSerializer struct{
	Owner string `json:"owner"`
	ListName string `json:"listName"`
	Title    string `json:"title"`
    Writer   string `json:"writer"`
    Remark   string `json:"remark"`
}

func createBookSerializer(book models.Book) BookSerializer{
	return BookSerializer{
		ListName: book.ListName,
		Title: book.Title,
		Writer: book.Writer,
		Remark: book.Remark,
		Owner: book.Owner,
	}
}

func AddBook(ctx *gin.Context) {
	owner := ctx.Query("owner")
	if owner == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"An activity owner cannot be empty",
		})	
	}

	var newBook models.Book

	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	newBook.Owner = owner
	err := database.DB.DB.Create(&newBook).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Could not add book to database.Try again later",
			"details":err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"created": true,
		"book": createBookSerializer(newBook),
		},
	)
}

func GetBook(ctx *gin.Context) {
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

	var books []models.Book
	var bookSerializers []BookSerializer
	err := database.DB.DB.Where(&models.Book{Owner: owner}).First(&books,"list_name = ?",title).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("The List %s could not be found", title),
		})
		return 
	}
	
	for _,b := range books{
		bookSerializers = append(bookSerializers, createBookSerializer(b))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": bookSerializers,
	})
}