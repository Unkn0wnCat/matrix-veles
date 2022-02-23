package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DBEntry struct {
	ID        primitive.ObjectID    `bson:"_id" json:"id"`
	Tags      []string              `bson:"tags" json:"tags"`           // Tags used for searching entries and ordering
	PartOf    []*primitive.ObjectID `bson:"part_of" json:"part_of"`     // PartOf specifies the lists this entry is part of
	HashValue string                `bson:"hash_value" json:"hash"`     // HashValue is the SHA512-hash of the file
	FileURL   string                `bson:"file_url" json:"file_url"`   // FileURL may be set to a file link
	Timestamp time.Time             `bson:"timestamp" json:"timestamp"` // Timestamp of when this entry was added
	AddedBy   *primitive.ObjectID   `bson:"added_by" json:"added_by"`   // AddedBy is a reference to the user who added this
	Comments  []*DBComment          `bson:"comments" json:"comments"`   // Comments regarding this entry
}
