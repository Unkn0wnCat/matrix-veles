package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DBEntry struct {
	ID        primitive.ObjectID    `bson:"_id" json:"id"`
	Tags      []string              `bson:"tags" json:"tags"`
	PartOf    []*primitive.ObjectID `bson:"part_of" json:"part_of"`
	HashValue string                `bson:"hash_value" json:"hash"`
	FileURL   string                `bson:"file_url" json:"file_url"`
	Timestamp time.Time             `bson:"timestamp" json:"timestamp"`
	AddedBy   *primitive.ObjectID   `bson:"added_by" json:"added_by"`
	Comments  []*DBComment          `bson:"comments" json:"comments"`
}
