package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	Model
	Name     string `json:"name"`
	Summary  string `json:"summary"`
	Author   Author `json:"author"`
	AuthorID int

	Categories []Category `json:"categories"`
}

func (m *Book) PreloadBook(db *gorm.DB) *gorm.DB {
	return db.Preload("Categories").Preload("Author")
}
