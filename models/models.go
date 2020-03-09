package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type WinningNumbers struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	WinningNumbers   string   `json:"winningnumbers,omitempty" bson:"winningnumbers,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}