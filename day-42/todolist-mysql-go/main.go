package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type task struct {
	Taskname string
	Complete bool
}
type Person struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Name        string
	Age         int
	Description string
	Work        []task
	CreatedAT   time.Time
	UpdatedAt   time.Time
}

var client *mongo.Client
var err error

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(r.Body)
	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)
	collection := client.Database("<dbname>").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println(person)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)
}

func fetchall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	collection := client.Database("<dbname>").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = result.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)
	json.NewEncoder(w).Encode(episodes)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}

// func updateUser(){

// }

// func getUserTask(){

// }

func main() {
	log.Info("starting server")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Ankita:Ank%401297@cluster0-evm4q.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	router := mux.NewRouter()
	router.HandleFunc("/add", createUser).Methods("POST")
	router.HandleFunc("/allUser", fetchall).Methods("GET")
	http.ListenAndServe(":8080", router)
}
