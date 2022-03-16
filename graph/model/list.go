package model

import (
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Tags          []string `json:"tags"`
	RawComments   []*model.DBComment
	MaintainerIDs []*primitive.ObjectID
}

func MakeList(dbList *model.DBHashList) *List {
	return &List{
		ID:            dbList.ID.Hex(),
		Name:          dbList.Name,
		Tags:          dbList.Tags,
		RawComments:   dbList.Comments,
		MaintainerIDs: dbList.Maintainers,
	}
}
