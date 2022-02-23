package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBComment struct {
	CommentedBy *primitive.ObjectID `bson:"commented_by" json:"commented_by"`
	Content     string              `bson:"content" json:"content"`
}
