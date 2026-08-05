package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	WordDb *mongo.Collection
	Hasii  *mongo.Collection
}

func ConnectDB(mongoURL string) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	// ping check
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	wordDb := client.Database("Word").Collection("WordDb")
	hasii := client.Database("HasiiDb").Collection("Hasii")

	return &Database{
		Client: client,
		WordDb: wordDb,
		Hasii:  hasii,
	}
}
