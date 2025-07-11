// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"petshop-backend/config"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // Models
// type Owner struct {
// 	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name  string             `json:"name" bson:"name"`
// 	Email string             `json:"email" bson:"email"`
// 	Phone string             `json:"phone" bson:"phone"`
// }

// type Pet struct {
// 	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name    string             `json:"name" bson:"name"`
// 	Species string             `json:"species" bson:"species"`
// 	Age     int                `json:"age" bson:"age"`
// 	OwnerID primitive.ObjectID `json:"owner_id" bson:"owner_id"`
// }

// type Service struct {
// 	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Name        string             `json:"name" bson:"name"`
// 	Description string             `json:"description" bson:"description"`
// 	Price       int                `json:"price" bson:"price"`
// }

// type Appointment struct {
// 	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	PetID     primitive.ObjectID `json:"pet_id" bson:"pet_id"`
// 	ServiceID primitive.ObjectID `json:"service_id" bson:"service_id"`
// 	Date      string             `json:"date" bson:"date"`
// 	Note      string             `json:"note" bson:"note"`
// }

// func main() {
// 	config.ConnectDB()

// 	// Create database and collections
// 	db := config.DB.Database("petshop")
// 	owners := db.Collection("owners")
// 	pets := db.Collection("pets")
// 	services := db.Collection("services")
// 	appointments := db.Collection("appointments")

// 	// Seed data
// 	seedOwners(owners, 20)
// 	seedPets(pets, 20)
// 	seedServices(services, 20)
// 	seedAppointments(appointments, 20)

// 	fmt.Println("Database seeding completed successfully!")
// }

// func seedOwners(collection *mongo.Collection, count int) {
// 	owners := make([]Owner, count)
// 	for i := 0; i < count; i++ {
// 		owners[i] = Owner{
// 			ID:    primitive.NewObjectID(),
// 			Name:  generateName(),
// 			Email: fmt.Sprintf("owner%d@example.com", i+1),
// 			Phone: generatePhone(),
// 		}
// 	}

// 	_, err := collection.InsertMany(context.TODO(), toInterfaceSlice(owners))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func seedPets(collection *mongo.Collection, count int) {
// 	owners, err := getOwners(collection)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pets := make([]Pet, count)
// 	species := []string{"Dog", "Cat", "Bird", "Fish", "Rabbit"}

// 	for i := 0; i < count; i++ {
// 		pets[i] = Pet{
// 			ID:      primitive.NewObjectID(),
// 			Name:    generatePetName(),
// 			Species: species[rand.Intn(len(species))],
// 			Age:     rand.Intn(10) + 1, // Age between 1 and 10
// 			OwnerID: owners[rand.Intn(len(owners))].ID,
// 		}
// 	}

// 	_, err = collection.InsertMany(context.TODO(), toInterfaceSlice(pets))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func seedServices(collection *mongo.Collection, count int) {
// 	services := []interface{}{
// 		Service{ID: primitive.NewObjectID(), Name: "Grooming", Description: "Full body grooming", Price: 150000},
// 		Service{ID: primitive.NewObjectID(), Name: "Vaccination", Description: "Annual vaccination", Price: 200000},
// 		Service{ID: primitive.NewObjectID(), Name: "Dental Cleaning", Description: "Professional dental cleaning", Price: 300000},
// 		Service{ID: primitive.NewObjectID(), Name: "Nail Trim", Description: "Professional nail trimming", Price: 50000},
// 		Service{ID: primitive.NewObjectID(), Name: "Bath", Description: "Professional bathing", Price: 100000},
// 		// Add more services...
// 	}

// 	_, err := collection.InsertMany(context.TODO(), services)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func seedAppointments(collection *mongo.Collection, count int) {
// 	pets, err := getPets(collection.Database().Collection("pets"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	services, err := getServices(collection.Database().Collection("services"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if len(pets) == 0 || len(services) == 0 {
// 		log.Fatal("Cannot seed appointments: no pets or services found. Please seed pets and services first.")
// 	}

// 	appointments := make([]Appointment, 0, count)
// 	for i := 0; i < count; i++ {
// 		appointments = append(appointments, Appointment{
// 			ID:        primitive.NewObjectID(),
// 			PetID:     pets[rand.Intn(len(pets))].ID,
// 			ServiceID: services[rand.Intn(len(services))].ID,
// 			Date:      generateAppointmentDate(),
// 			Note:      generateAppointmentNote(),
// 		})
// 	}

// 	if len(appointments) > 0 {
// 		_, err = collection.InsertMany(context.TODO(), toInterfaceSlice(appointments))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func generateName() string {
// 	firstNames := []string{"John", "Sarah", "Michael", "Emily", "David", "Jessica", "Daniel", "Olivia", "James", "Sophia"}
// 	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}
// 	return fmt.Sprintf("%s %s", firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))])
// }

// func generatePhone() string {
// 	return fmt.Sprintf("+628%09d", rand.Int63n(1000000000))
// }

// func generatePetName() string {
// 	petNames := []string{"Bella", "Max", "Luna", "Charlie", "Oliver", "Lucy", "Cooper", "Daisy", "Rocky", "Milo"}
// 	return petNames[rand.Intn(len(petNames))]
// }

// func generateAppointmentDate() string {
// 	// Generate dates between now and 3 months from now
// 	now := time.Now()
// 	max := now.AddDate(0, 3, 0)
// 	timestamp := rand.Int63n(max.Unix()-now.Unix()) + now.Unix()
// 	return time.Unix(timestamp, 0).Format(time.RFC3339)
// }

// func generateAppointmentNote() string {
// 	notes := []string{
// 		"Regular checkup",
// 		"Annual vaccination",
// 		"First grooming session",
// 		"Follow-up appointment",
// 		"Emergency visit",
// 	}
// 	return notes[rand.Intn(len(notes))]
// }

// func toInterfaceSlice[T any](slice []T) []interface{} {
// 	result := make([]interface{}, len(slice))
// 	for i, v := range slice {
// 		result[i] = v
// 	}
// 	return result
// }

// func getOwners(collection *mongo.Collection) ([]Owner, error) {
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	var owners []Owner
// 	if err = cursor.All(context.TODO(), &owners); err != nil {
// 		return nil, err
// 	}
// 	return owners, nil
// }

// func getPets(collection *mongo.Collection) ([]Pet, error) {
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	var pets []Pet
// 	if err = cursor.All(context.TODO(), &pets); err != nil {
// 		return nil, err
// 	}
// 	return pets, nil
// }

// func getServices(collection *mongo.Collection) ([]Service, error) {
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	var services []Service
// 	if err = cursor.All(context.TODO(), &services); err != nil {
// 		return nil, err
// 	}
// 	return services, nil
// }
