package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBHashList struct {
	ID          primitive.ObjectID    `bson:"_id" json:"id"`
	Name        string                `bson:"name" json:"name"`
	Tags        []string              `bson:"tags" json:"tags"`               // Tags of this list for discovery, and sorting
	Comments    []*DBComment          `bson:"comments" json:"comments"`       // Comments regarding this list
	Maintainers []*primitive.ObjectID `bson:"maintainers" json:"maintainers"` // Maintainers contains references to the users who may edit this list
}
