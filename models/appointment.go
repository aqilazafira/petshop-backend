package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	PetID     primitive.ObjectID `json:"pet_id" bson:"pet_id"`
	ServiceID primitive.ObjectID `json:"service_id" bson:"service_id"`
	Date      string             `json:"date" bson:"date"`
	Note      string             `json:"note" bson:"note"`
}
