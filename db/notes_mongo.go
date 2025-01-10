package db

import (
	"context"
	"fmt"
	"noteservice/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongo() {

	clientOption := options.Client().ApplyURI("")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cl, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println("err", err)
	}

	client = cl
}

func InsertNote(notereq models.Note) {
	coll := client.Database("admindb").Collection("notes")
	// fmt.Println("Connected to Mongo DB!")

	note := bson.M{
		"title":   notereq.Title,
		"content": notereq.Content,
		"created": notereq.Created,
	}

	insertResult, err := coll.InsertOne(context.Background(), note)
	if err != nil {
		fmt.Println("insert err", err)
	}

	fmt.Println("Inserted with id", insertResult)
}

func ReadNotes() []models.Note {
	coll := client.Database("admindb").Collection("notes")
	// fmt.Println("Connected to Mongo DB!")

	var notes []models.Note

	cur, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("find err", err)
	}

	defer cur.Close(context.Background())

	cur.All(context.Background(), &notes)

	err = cur.Err()
	if err != nil {
		fmt.Println("cursor err", err)
	}

	return notes
}
