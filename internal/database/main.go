package internal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to the database")
		return nil

	}
	return db

}

// func (n *NotesService) FindNoteByTitle(title string) (*internal.Notes, error) {
// 	var note internal.Notes
// 	// Use GORM's First method to find a note by title
// 	result := n.db.Where("title = ?", title).First(&note)

// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, nil // No note found with the given title
// 		}
// 		return nil, result.Error // Unexpected error
// 	}
// 	return &note, nil
// }
