package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type DB struct {
	*mongo.Client
}

func NewDB(client *mongo.Client) *DB {
	return &DB{client}
}

func (db *DB) Save(guid string, token []byte) error {
	collection := db.Database("auth").Collection("users")
	dbDoc := bson.M{"guid": guid, "refreshToken": token}
	_, err := collection.InsertOne(context.Background(), dbDoc)
	return err
}
