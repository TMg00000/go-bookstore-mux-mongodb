package mongodb

import (
	"context"
	"go-bookstore-mux-mongodb/internal/configs"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoConnection(ctx context.Context) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(configs.Env.MongoURI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
