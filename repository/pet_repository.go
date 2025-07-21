package repository

import (
	"context"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPets() ([]models.Pet, error) {
	collection := config.DB.Database("petshop").Collection("pets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var pets []models.Pet
	if err = cursor.All(ctx, &pets); err != nil {
		return nil, err
	}

	return pets, nil
}

func GetPetByID(id primitive.ObjectID) (models.Pet, error) {
	collection := config.DB.Database("petshop").Collection("pets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var pet models.Pet
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&pet)
	return pet, err
}

func CreatePet(pet models.Pet) error {
	collection := config.DB.Database("petshop").Collection("pets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, pet)
	return err
}

func UpdatePet(id primitive.ObjectID, update bson.M) error {
	collection := config.DB.Database("petshop").Collection("pets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeletePet(id primitive.ObjectID) error {
	collection := config.DB.Database("petshop").Collection("pets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
