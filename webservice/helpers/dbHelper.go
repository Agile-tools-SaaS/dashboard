package helpers

import (
	"context"
	"log"
	"time"

	"github.com/Agile-tools-SaaS/dashboard/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(mongo_uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	return client
}

func newRepository(table string) *models.Repository {
	repository := new(models.Repository)
	repository.table = table
	return repository
}

func NewContext(table string) *mongo.Collection {
	repo := newRepository(table)
	return repo.NewConnection()
}
