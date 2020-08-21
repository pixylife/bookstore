package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	Model
	Type   string `json:"type"`
	Book   Book   `gorm:"foreignkey:bookID" json:"book"`
	BookID int
}

func (m *Category) PreloadCategory(db *gorm.DB) *gorm.DB {
	return db.Preload("Book")
}
