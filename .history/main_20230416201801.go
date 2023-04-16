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

func Connect(dbUrl string) *SB