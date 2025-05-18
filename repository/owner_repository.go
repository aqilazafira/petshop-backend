package repository

import (
	"context"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOwners() ([]models.Owner, error) {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var owners []models.Owner
	if err = cursor.All(ctx, &owners); err != nil {
		return nil, err
	}

	return owners, nil
}

func GetOwnerByID(id primitive.ObjectID) (models.Owner, error) {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var owner models.Owner
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&owner)
	return owner, err
}

func CreateOwner(owner models.Owner) error {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, owner)
	return err
}

func UpdateOwner(id primitive.ObjectID, update bson.M) error {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeleteOwner(id primitive.ObjectID) error {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func OwnerExistsByEmail(email string) (bool, error) {
	collection := config.DB.Database("petshop").Collection("owners")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{"email": email})
	return count > 0, err
}