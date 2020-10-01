package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenPair struct {
	AccessToken, RefreshToken string
}

type Document struct {
	Guid         string `json:"guid,omitempty" bson:"guid,omitempty"`
	RefreshToken []byte `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

type DB struct {
	*mongo.Client
}

func NewDB(client *mongo.Client) *DB {
	return &DB{client}
}

func (db *DB) Save(guid string, token []byte) error {
	collection := db.Database("auth").Collection("users")
	dbDoc := Document{guid, token} // check
	_, err := collection.InsertOne(context.Background(), dbDoc)
	return err
}

func (db *DB) RemoveTokenFromDocument(guid string) error {
	collection := db.Database("auth").Collection("users")
	filter := bson.D{{"guid", guid}}
	update := bson.M{"$unset": bson.M{"refreshToken": ""}}
	_, err := collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (db *DB) FindOne(guid string) Document {
	var doc Document
	collection := db.Database("auth").Collection("users")
	res := collection.FindOne(context.TODO(), bson.D{{"guid", guid}})
	res.Decode(&doc)
	return doc
}

func (db *DB) IsExists(guid string) bool {
	collection := db.Database("auth").Collection("users")
	count, _ := collection.CountDocuments(context.TODO(), bson.M{"guid": guid})
	return count != 0
}
