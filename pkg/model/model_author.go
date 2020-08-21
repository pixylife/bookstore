package model

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	Model
	Name string `json:"name"`
}

func (m *Author) PreloadAuthor(db *gorm.DB) *gorm.DB {
	return db
}
