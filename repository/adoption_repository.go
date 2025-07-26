package repository

import (
	"context"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAdoptions() ([]models.Adoption, error) {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var adoptions []models.Adoption
	if err = cursor.All(ctx, &adoptions); err != nil {
		return nil, err
	}

	return adoptions, nil
}

func GetAdoptionByID(id primitive.ObjectID) (models.Adoption, error) {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var adoption models.Adoption
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&adoption)
	return adoption, err
}

func CreateAdoption(adoption models.Adoption) error {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, adoption)
	return err
}

func UpdateAdoption(id primitive.ObjectID, update bson.M) error {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeleteAdoption(id primitive.ObjectID) error {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func GetAdoptionsByStatus(status string) ([]models.Adoption, error) {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"status": status})
	if err != nil {
		return nil, err
	}

	var adoptions []models.Adoption
	if err = cursor.All(ctx, &adoptions); err != nil {
		return nil, err
	}

	return adoptions, nil
}

func GetAdoptionsByPetID(petID primitive.ObjectID) ([]models.Adoption, error) {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"pet_id": petID})
	if err != nil {
		return nil, err
	}

	var adoptions []models.Adoption
	if err = cursor.All(ctx, &adoptions); err != nil {
		return nil, err
	}

	return adoptions, nil
}

func GetAdoptionsByUserEmail(userEmail string) ([]models.Adoption, error) {
	collection := config.DB.Database("petshop").Collection("adoptions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"user_email": userEmail})
	if err != nil {
		return nil, err
	}

	var adoptions []models.Adoption
	if err = cursor.All(ctx, &adoptions); err != nil {
		return nil, err
	}

	return adoptions, nil
}
