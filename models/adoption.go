package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Adoption struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PetID            primitive.ObjectID `json:"pet_id" bson:"pet_id"`
	PetName          string             `json:"pet_name" bson:"pet_name"`
	Name             string             `json:"name" bson:"name"`
	Email            string             `json:"email" bson:"email"`
	Phone            string             `json:"phone" bson:"phone"`
	Address          string             `json:"address" bson:"address"`
	Experience       string             `json:"experience" bson:"experience"`
	Reason           string             `json:"reason" bson:"reason"`
	LivingSpace      string             `json:"living_space" bson:"living_space"`
	HasOtherPets     bool               `json:"has_other_pets" bson:"has_other_pets"`
	OtherPetsDetails string             `json:"other_pets_details" bson:"other_pets_details"`
	Status           string             `json:"status" bson:"status"` // pending, approved, rejected
	SubmissionDate   time.Time          `json:"submission_date" bson:"submission_date"`
	AdoptionDate     *time.Time         `json:"adoption_date,omitempty" bson:"adoption_date,omitempty"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}
