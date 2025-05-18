package repository

import (
	"context"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAppointments() ([]models.Appointment, error) {
	collection := config.DB.Database("petshop").Collection("appointmnets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	if err = cursor.All(ctx, &appointments); err != nil {
		return nil, err
	}

	return appointments, nil
}

func GetAppointmentByID(id primitive.ObjectID) (models.Appointment, error) {
	collection := config.DB.Database("petshop").Collection("appointmnets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var appointment models.Appointment
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&appointment)
	return appointment, err
}

func CreateAppointment(appointment models.Appointment) error {
	collection := config.DB.Database("petshop").Collection("appointmnets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, appointment)
	return err
}

func UpdateAppointment(id primitive.ObjectID, update bson.M) error {
	collection := config.DB.Database("petshop").Collection("appointmnets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeleteAppointment(id primitive.ObjectID) error {
	collection := config.DB.Database("petshop").Collection("appointmnets")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}