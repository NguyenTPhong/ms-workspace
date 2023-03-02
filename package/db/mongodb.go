package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(connStr, dbName string) (*mongo.Database, func(), error) {

	structCodec, _ := bsoncodec.NewStructCodec(bsoncodec.JSONFallbackStructTagParser)
	rb := bson.NewRegistryBuilder()
	rb.RegisterDefaultEncoder(reflect.Struct, structCodec)
	rb.RegisterDefaultDecoder(reflect.Struct, structCodec)
	clientOptions := options.Client().SetRegistry(rb.Build()).ApplyURI(connStr)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, func() {}, err
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, func() {}, err
	}

	db := client.Database(dbName)
	return db, func() {
		err = client.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}, nil
}
