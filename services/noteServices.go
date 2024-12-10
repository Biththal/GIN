package services

import (
	internal "GIN/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {

	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Biththal",
		},
		{
			Id:   2,
			Name: "Success",
		},
	}
	return data
}

func (n *NotesService) CreatePostService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	err := n.db.Create(note).Error

	if err != nil {
		fmt.Print(err)
		return nil, err

	}

	return note, nil

}
