package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBHashList struct {
	ID          primitive.ObjectID    `bson:"_id" json:"id"`
	Tags        []string              `bson:"tags" json:"tags"`
	Comments    []*DBComment          `bson:"comments" json:"comments"`
	Maintainers []*primitive.ObjectID `bson:"maintainers" json:"maintainers"`
}
