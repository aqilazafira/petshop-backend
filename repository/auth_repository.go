package repository

import (
	"context"
	"fmt"
	"petshop-backend/config"
	"petshop-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByEmail(ctx context.Context, email string) (*models.UserLogin, error) {
	collection := config.DB.Database("petshop").Collection("users")

	var user models.UserLogin
	filter := bson.M{"email": email}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("email %s tidak ditemukan", email)
		}
		return nil, err
	}

	return &user, nil
}

func InsertUser(ctx context.Context, user models.UserLogin) (interface{}, error) {
	collection := config.DB.Database("petshop").Collection("users")

	// Cek apakah email sudah ada
	filter := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("email %s sudah digunakan", user.Email)
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}
