package controllers

import (
	"GIN/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {

	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreatePost())
	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}

}
func (n *NotesController) CreatePost() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}
	return func(ctx *gin.Context) {

		var noteBody NoteBody
		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"Message": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreatePostService(noteBody.Title, noteBody.Status)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
			return

		}

		ctx.JSON(200, gin.H{
			"notes": note,
		})
	}

}
