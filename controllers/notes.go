package controllers

import "github.com/gin-gonic/gin"

type NotesController struct{}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {

	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreatePost())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Get Request",
		})
	}

}
func (n *NotesController) CreatePost() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": "Post Request",
		})
	}

}
