package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SpaceFile struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string
	Content string
}
