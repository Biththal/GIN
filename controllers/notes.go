package controllers

import (
	"GIN/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {

	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreatePost())
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())

	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		status := ctx.Query("status")
		actualstatus, err := strconv.ParseBool(status)
		if err != nil {
			fmt.Print(err)
		}
		notes, err := n.notesService.GetNotesService(actualstatus)
		if err != nil {

			ctx.JSON(400, gin.H{
				"Message": err.Error(),
			})

		}
		ctx.JSON(200, gin.H{
			"notes": notes,
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
func (n *NotesController) UpdateNotes() gin.HandlerFunc {

	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {

		var noteBody NoteBody
		if err := ctx.BindJSON(&noteBody); err != nil {
			ctx.JSON(400, gin.H{
				"Message": err.Error(),
			})
			return
		}
		note, err := n.notesService.UpdateNotesService(noteBody.Title, noteBody.Status, noteBody.Id)
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
func (n *NotesController) DeleteNotes() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		noteID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err,
			})
		}

		err = n.notesService.DeleteNotesServices(noteID)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return

		}

		ctx.JSON(200, gin.H{
			"messege": "Note deleted successfully",
		})

	}
}
