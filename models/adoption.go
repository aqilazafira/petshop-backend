package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Adoption struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	PetID        primitive.ObjectID `json:"pet_id" bson:"pet_id"`
	OwnerID      primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	AdoptionDate primitive.DateTime `json:"adoption_date" bson:"adoption_date"`
	Status       string             `json:"status" bson:"status"`
}
