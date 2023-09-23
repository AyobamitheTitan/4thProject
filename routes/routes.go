package routes

import (
	"v1/listee/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRoutes(engine *gin.Engine) {
	// user routes
	user := engine.Group("/user")
	{
		user.POST("/signup",controllers.Signup)
		user.POST("/login",controllers.Login)
	}

	// book routes
	books := engine.Group("/books")
	{
		books.GET("/",controllers.GetBook)	
		books.POST("/",controllers.AddBook)
	}

	// song routes
	songs := engine.Group("/songs")
	{
		songs.GET("/",controllers.GetSong)
		songs.POST("/",controllers.AddSong)
	}

	// movies routes
	movies := engine.Group("/movies")
	{
		movies.GET("/",controllers.GetMovie)
		movies.POST("/",controllers.AddMovie)
	}

	// others routes
	others := engine.Group("/others")
	{
		others.GET("/",controllers.GetOther)
		others.POST("/",controllers.AddOther)
	}

	everything := engine.Group("/list")
	{
		everything.GET("/",controllers.GetEverything)
	}
}

