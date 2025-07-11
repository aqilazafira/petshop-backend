package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Owner struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	Phone   string             `json:"phone" bson:"phone"`
	Address string             `json:"address" bson:"address"`
}
