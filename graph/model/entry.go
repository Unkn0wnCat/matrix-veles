package model

import (
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Entry struct {
	ID          string   `json:"id"`
	Tags        []string `json:"tags"`
	PartOfIDs   []*primitive.ObjectID
	HashValue   string    `json:"hashValue"`
	FileURL     *string   `json:"fileUrl"`
	Timestamp   time.Time `json:"timestamp"`
	AddedByID   primitive.ObjectID
	RawComments []*model.DBComment
}

func MakeEntry(dbEntry *model.DBEntry) *Entry {
	return &Entry{
		ID:        dbEntry.ID.Hex(),
		Tags:      dbEntry.Tags,
		PartOfIDs: dbEntry.PartOf,
		HashValue: dbEntry.HashValue,
		//FileURL:     &dbEntry.FileURL,
		Timestamp:   dbEntry.Timestamp,
		AddedByID:   *dbEntry.AddedBy,
		RawComments: dbEntry.Comments,
	}
}
