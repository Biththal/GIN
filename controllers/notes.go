package controllers

import (
	"GIN/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {

	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreatePost())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}

}
func (n *NotesController) CreatePost() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.CreatePostService(),
		})
	}

}
