package main

import (
	"GIN/controllers"

	"GIN/services"

	internal "GIN/internal/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := internal.InitDB()

	if db == nil {
		fmt.Print("error while connecting to the database")
	} else {
		fmt.Println("Database connected successfully!!!")
	}
	notesService := &services.NotesService{}
	notesService.InitService(db)
	notesContrroler := &controllers.NotesController{}
	notesContrroler.InitNotesControllerRoutes(r, *notesService)

	r.Run(":8000")

	//r.GET("/ping", func(c *gin.Context) { //http://localhost:8000/ping
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong from main ",
	// })
	// })

	// r.GET("/hello", func(c *gin.Context) { //http://localhost:8000/hello
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })
	// r.GET("/me/:id", func(ctx *gin.Context) {
	// 	var id = ctx.Param("id")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"user_id": id,
	// 	})

	// })
	// r.POST("/submit", func(c *gin.Context) { ////http://localhost:8000/submit

	// 	name := c.PostForm("name")
	// 	//text := c.GetString("hello world")
	// 	c.JSON(200, gin.H{

	// 		"text": name,
	// 	})
	// })

	// r.POST("/new", func(ctx *gin.Context) {
	// 	type myRequest struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}

	// 	var MyRequest myRequest
	// 	ctx.BindJSON(&MyRequest)
	// 	ctx.JSON(200, gin.H{

	// 		"email":    MyRequest.Email,
	// 		"password": MyRequest.Password,
	// 	})
	// })
	// r.PUT("/new", func(ctx *gin.Context) {
	// 	type myRequest struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}

	// 	var MyRequest myRequest
	// 	ctx.BindJSON(&MyRequest)
	// 	ctx.JSON(200, gin.H{

	// 		"email":    MyRequest.Email,
	// 		"password": MyRequest.Password,
	// 	})
	// })

	// r.PATCH("/new", func(ctx *gin.Context) {
	// 	type myRequest struct {
	// 		Email    string `json:"email"`
	// 		Password string `json:"password"`
	// 	}

	// 	var MyRequest myRequest
	// 	ctx.BindJSON(&MyRequest)
	// 	ctx.JSON(200, gin.H{

	// 		"email":    MyRequest.Email,
	// 		"password": MyRequest.Password,
	// 	})
	// })

	// r.DELETE("me/:id", func(ctx *gin.Context) {
	// 	var id = ctx.Param("id")

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"id":      id,
	// 		"message": "Deleted!!!",
	// 	})
	// })

}
