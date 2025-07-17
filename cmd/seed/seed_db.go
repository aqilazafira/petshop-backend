package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"petshop-backend/config"
	"petshop-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config.ConnectDB()

	// Create database and collections
	db := config.DB.Database("petshop")
	owners := db.Collection("owners")
	pets := db.Collection("pets")
	services := db.Collection("services")
	appointments := db.Collection("appointments")

	// Seed data
	seedOwners(owners, 20)
	seedPets(pets, 20)
	seedServices(services)
	seedAppointments(appointments, 20)

	fmt.Println("Database seeding completed successfully!")
}

func seedOwners(collection *mongo.Collection, count int) {
	owners := make([]models.Owner, count)
	for i := 0; i < count; i++ {
		owners[i] = models.Owner{
			ID:      primitive.NewObjectID(),
			Name:    generateName(),
			Email:   fmt.Sprintf("owner%d@example.com", i+1),
			Phone:   generatePhone(),
			Address: generateAddress(),
		}
	}

	_, err := collection.InsertMany(context.TODO(), toInterfaceSlice(owners))
	if err != nil {
		log.Fatal(err)
	}
}

func seedPets(collection *mongo.Collection, count int) {
	owners, err := getOwners(collection.Database().Collection("owners"))
	if err != nil {
		log.Fatal(err)
	}

	pets := make([]models.Pet, count)
	species := []string{"Dog", "Cat", "Bird", "Fish", "Rabbit"}
	genders := []string{"Male", "Female"}

	for i := 0; i < count; i++ {
		pets[i] = models.Pet{
			ID:       primitive.NewObjectID(),
			Name:     generatePetName(),
			Species:  species[rand.Intn(len(species))],
			Age:      rand.Intn(10) + 1, // Age between 1 and 10
			OwnerID:  owners[rand.Intn(len(owners))].ID,
			Gender:   genders[rand.Intn(len(genders))],
			ImageURL: generatePetImageURL(),
		}
	}

	_, err = collection.InsertMany(context.TODO(), toInterfaceSlice(pets))
	if err != nil {
		log.Fatal(err)
	}
}

func seedServices(collection *mongo.Collection) {
	services := []interface{}{
		models.Service{ID: primitive.NewObjectID(), Name: "Grooming", Description: "Full body grooming for your beloved pet.", Price: 150000},
		models.Service{ID: primitive.NewObjectID(), Name: "Vaccination", Description: "Complete annual vaccination package.", Price: 200000},
		models.Service{ID: primitive.NewObjectID(), Name: "Dental Cleaning", Description: "Professional dental cleaning and polishing.", Price: 300000},
		models.Service{ID: primitive.NewObjectID(), Name: "Nail Trim", Description: "Gentle and precise nail trimming service.", Price: 50000},
		models.Service{ID: primitive.NewObjectID(), Name: "Bath & Brush", Description: "A refreshing bath and thorough brushing.", Price: 100000},
		models.Service{ID: primitive.NewObjectID(), Name: "Parasite Control", Description: "Effective flea, tick, and mite treatment.", Price: 120000},
		models.Service{ID: primitive.NewObjectID(), Name: "Microchipping", Description: "Permanent identification for your pet's safety.", Price: 250000},
		models.Service{ID: primitive.NewObjectID(), Name: "Health Check-up", Description: "Comprehensive health examination by our vet.", Price: 180000},
	}

	_, err := collection.InsertMany(context.TODO(), services)
	if err != nil {
		log.Fatal(err)
	}
}

func seedAppointments(collection *mongo.Collection, count int) {
	pets, err := getPets(collection.Database().Collection("pets"))
	if err != nil {
		log.Fatal(err)
	}

	services, err := getServices(collection.Database().Collection("services"))
	if err != nil {
		log.Fatal(err)
	}

	if len(pets) == 0 || len(services) == 0 {
		log.Fatal("Cannot seed appointments: no pets or services found. Please seed pets and services first.")
	}

	appointments := make([]models.Appointment, 0, count)
	for i := 0; i < count; i++ {
		appointments = append(appointments, models.Appointment{
			ID:        primitive.NewObjectID(),
			PetID:     pets[rand.Intn(len(pets))].ID,
			ServiceID: services[rand.Intn(len(services))].ID,
			Date:      generateAppointmentDate(),
			Note:      generateAppointmentNote(),
		})
	}

	if len(appointments) > 0 {
		_, err = collection.InsertMany(context.TODO(), toInterfaceSlice(appointments))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func generateName() string {
	firstNames := []string{"John", "Sarah", "Michael", "Emily", "David", "Jessica", "Daniel", "Olivia", "James", "Sophia"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}
	return fmt.Sprintf("%s %s", firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))])
}

func generatePhone() string {
	return fmt.Sprintf("+628%09d", rand.Int63n(1000000000))
}

func generateAddress() string {
	streets := []string{"Jl. Sudirman", "Jl. Thamrin", "Jl. Gatot Subroto", "Jl. Rasuna Said", "Jl. Diponegoro"}
	cities := []string{"Jakarta", "Bandung", "Surabaya", "Medan", "Makassar"}
	return fmt.Sprintf("%s No. %d, %s", streets[rand.Intn(len(streets))], rand.Intn(100)+1, cities[rand.Intn(len(cities))])
}

func generatePetName() string {
	petNames := []string{"Bella", "Max", "Luna", "Charlie", "Oliver", "Lucy", "Cooper", "Daisy", "Rocky", "Milo"}
	return petNames[rand.Intn(len(petNames))]
}

func generatePetImageURL() string {
	images := []string{
		"/pets/Aqila.png",
	}
	return images[rand.Intn(len(images))]
}

func generateAppointmentDate() string {
	// Generate dates between now and 3 months from now
	now := time.Now()
	max := now.AddDate(0, 3, 0)
	timestamp := rand.Int63n(max.Unix()-now.Unix()) + now.Unix()
	return time.Unix(timestamp, 0).Format(time.RFC3339)
}

func generateAppointmentNote() string {
	notes := []string{
		"Regular checkup",
		"Annual vaccination",
		"First grooming session",
		"Follow-up appointment",
		"Emergency visit",
		"Needs special attention",
		"Allergic to certain foods",
		"Very friendly with other pets",
	}
	return notes[rand.Intn(len(notes))]
}

func toInterfaceSlice[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

func getOwners(collection *mongo.Collection) ([]models.Owner, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var owners []models.Owner
	if err = cursor.All(context.TODO(), &owners); err != nil {
		return nil, err
	}
	return owners, nil
}

func getPets(collection *mongo.Collection) ([]models.Pet, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var pets []models.Pet
	if err = cursor.All(context.TODO(), &pets); err != nil {
		return nil, err
	}
	return pets, nil
}

func getServices(collection *mongo.Collection) ([]models.Service, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var services []models.Service
	if err = cursor.All(context.TODO(), &services); err != nil {
		return nil, err
	}
	return services, nil
}
