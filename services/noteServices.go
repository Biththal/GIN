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

func (n *NotesService) GetNotesService(status bool) ([]*internal.Notes, error) {

	var notes []*internal.Notes
	if err := n.db.Where("status= ?", status).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NotesService) CreatePostService(title string, status bool) (*internal.Notes, error) {

	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	// if note.Title == "" {
	// 	return nil, errors.New("Title is required")

	// }

	err := n.db.Create(note).Error

	if err != nil {
		fmt.Print(err)
		return nil, err

	}

	return note, nil

}
func (n *NotesService) UpdateNotesService(title string, status bool, id int) (*internal.Notes, error) {

	var note *internal.Notes
	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	err := n.db.Save(note).Error

	if err != nil {
		fmt.Print(err)
		return nil, err

	}

	return note, nil

}

func (n *NotesService) DeleteNotesServices(id int64) error {

	var note *internal.Notes
	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}
