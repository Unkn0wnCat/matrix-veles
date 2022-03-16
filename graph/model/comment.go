package model

import (
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	Timestamp time.Time `json:"timestamp"`
	AuthorID  *primitive.ObjectID
	Content   string `json:"content"`
}

func MakeComment(dbComment *model.DBComment) *Comment {
	return &Comment{
		Timestamp: dbComment.Timestamp,
		AuthorID:  dbComment.CommentedBy,
		Content:   dbComment.Content,
	}
}
