package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DBComment struct {
	Timestamp   time.Time           `bson:"timestamp" json:"timestamp"`
	CommentedBy *primitive.ObjectID `bson:"commented_by" json:"commented_by"` // CommentedBy contains a reference to the user who commented
	Content     string              `bson:"content" json:"content"`           // Content is the body of the comment
}
