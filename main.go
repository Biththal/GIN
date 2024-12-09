package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { //http://localhost:8000/ping
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from main ",
		})
	})

	r.GET("/hello", func(c *gin.Context) { //http://localhost:8000/hello
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/me/:id", func(ctx *gin.Context) {
		var id = ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})

	})
	r.POST("/submit", func(c *gin.Context) { ////http://localhost:8000/submit

		name := c.PostForm("name")
		//text := c.GetString("hello world")
		c.JSON(200, gin.H{

			"text": name,
		})
	})

	r.POST("/new", func(ctx *gin.Context) {
		type myRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var MyRequest myRequest
		ctx.BindJSON(&MyRequest)
		ctx.JSON(200, gin.H{

			"email":    MyRequest.Email,
			"password": MyRequest.Password,
		})
	})
	r.Run(":8000")
}
