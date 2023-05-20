package helpers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(mongo_uri string) (*mongo.Client, *context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// defer client.Disconnect(ctx)

	return client, &ctx
}

type Context struct {
	// db         *mongo.Client
	Context    *context.Context
	Collection *mongo.Collection
	Close      func()
}

func NewContext(table_name string) *Context {
	mongo_uri := GetEnvByName("MONGO_URI")
	db, db_ctx := InitDB(mongo_uri)

	ctx := Context{
		Context:    db_ctx,
		Close:      func() { db.Disconnect(*db_ctx) },
		Collection: db.Database("ASCore").Collection(table_name),
	}

	return &ctx
}
