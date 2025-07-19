package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       int                `json:"price" bson:"price"`
}
