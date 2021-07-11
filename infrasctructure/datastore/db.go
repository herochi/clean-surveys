package datastore

import (
	"context"

	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Database, error) {
	ENV := viper.GetString("ENV")
	uri := viper.GetString("MONGO.URI")
	if ENV == "DEV" {
		uri = viper.GetString("MONGO.URI_DEV")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString("MONGO.DATABASE")), nil
}
