package mongo

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(url string) (*mongo.Client, error) {
	slog.Info("Connecting to Mongo...")

	clOpts := options.Client().ApplyURI(url)

	cl, err := mongo.Connect(context.TODO(), clOpts)
	if err != nil {
		return nil, err
	}

	if err := cl.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	slog.Info("Connected to Mongo successfully!")

	return cl, nil
}
