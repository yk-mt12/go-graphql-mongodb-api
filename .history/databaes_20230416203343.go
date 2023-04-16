package database

import (
	"time"
	"fmt"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo"
	"go-graphql-mongodb-api/graph/model"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB struct {
	client *mongo.Client
}

// Connect to MongoDB
func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &DB{client: client}

}

// 
func (db *DB) InsertMovieById(movie model.NewMovie) *model.Movie {
	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := movieColl.InsertOne(ctx, bson.D{{Key: "name", Value: movie.Name}})

	if err != nil {
			log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnMovie := model.Movie{ID: insertedID, Name: movie.Name}

	return &returnMovie
}