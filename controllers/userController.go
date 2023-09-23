package controllers

import (
	"net/http"
	"v1/listee/database"
	"v1/listee/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type userSerializer struct {
// 	Username string `json:"username"`
// 	FirstName string `json:"firstName"`
// 	LastName string `json:"lastName"`
// 	Password string `json:"password"`
// 	Lists []string `json:"lists"`
// }

type loginSerialzer struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(ctx *gin.Context){
	var newUser models.User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Bad request",
			"details":err.Error(),
		})
		return
	}

	hashed,err := bcrypt.GenerateFromPassword([]byte(newUser.Password),14)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Unable to hash the password",
			"details":err.Error(),
		})
	}

	newUser.Password = string(hashed)

	err = database.DB.DB.Create(&newUser).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"Unable to add new user to the database",
			"details":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"created":true,
	})
}

func Login(ctx *gin.Context) {
	var user models.User
	var login loginSerialzer

	if err := ctx.BindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Bad request",
			"details":err.Error(),
		})
		return
	}

	err := database.DB.DB.Find(&user,"username = ?",login.Username).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error":"An unexpected error occurred. Please try again.",
			"details":err.Error(),
		})
	}


	if user.Username == "" {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error":"This user does not exist in the database",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Incorrect Password",
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"found":true,
	})
}