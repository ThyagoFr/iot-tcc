package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDBClient(conf *Config) (*mongo.Database,error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DNS()))
	if err != nil {
		return nil,err
	}
	return client.Database(conf.Database),nil
}