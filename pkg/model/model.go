package model

import (
	"fmt"
	"time"
)

import (
	"bookstore/pkg/eventing"
	"context"
	"github.com/jinzhu/gorm"
)

type Projection struct {
	eventing.BasicProjection `bson:",inline"`
	Data                     *Book `json:"Book" bson:"data"`
}
type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func BuildProjection(ctx context.Context, events []eventing.Event) (eventing.Projection, error) {
	ex := &Projection{}
	for _, event := range events {
		ex.Apply(event)
	}
	return ex, nil
}

func (ex *Projection) Apply(event eventing.Event) {

	ex.BasicProjection.Apply(event)
	switch v := event.(type) {

	case *Bookcreated:
		ex.Data = v.Data
	default:
		fmt.Println("could not handle event")
	}
}
func InitModels(db *gorm.DB) {

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Category{})
}
