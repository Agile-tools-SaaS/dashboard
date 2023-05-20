package models

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	table string
}

func (r *Repository) NewConnection() *mongo.Collection {
	mongo_uri := helpers.GetEnvByName("MONGO_URI")
	db := helpers.InitDB(mongo_uri)

	return db.Database("ASCore").Collection(r.table)
}
