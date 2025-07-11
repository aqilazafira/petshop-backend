package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pet struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerID  primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Name     string             `json:"name" bson:"name"`
	Species  string             `json:"species" bson:"species"`
	Age      int                `json:"age" bson:"age"`
	Gender   string             `json:"gender" bson:"gender"`
	ImageURL string             `json:"image_url" bson:"image_url"`
}
