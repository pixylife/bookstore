package model

import (
	"bookstore/pkg/eventing"
)

type Bookcreated struct {
	eventing.BasicEvent `bson:",inline"`
	Data                *Book `json:"data" bson:"data"`
}

func BookcreatedEvent(data *Book) *Bookcreated {

	return &Bookcreated{

		BasicEvent: eventing.NewBasicEvent(string(data.ID), "Bookcreated"),
		Data:       data,
	}
}
