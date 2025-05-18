package repository

import (
	"context"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetServices() ([]models.Service, error) {
	collection := config.DB.Database("petshop").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var services []models.Service
	if err = cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func GetServiceByID(id primitive.ObjectID) (models.Service, error) {
	collection := config.DB.Database("petshop").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var service models.Service
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&service)
	return service, err
}

func CreateService(service models.Service) error {
	collection := config.DB.Database("petshop").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, service)
	return err
}

func UpdateService(id primitive.ObjectID, update bson.M) error {
	collection := config.DB.Database("petshop").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeleteService(id primitive.ObjectID) error {
	collection := config.DB.Database("petshop").Collection("services")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
