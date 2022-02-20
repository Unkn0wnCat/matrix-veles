package db

import (
	"context"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DbClient *mongo.Client

func Connect() {
	if viper.GetString("bot.mongo.uri") == "" {
		log.Println("Skipping database login...")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newClient, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("bot.mongo.uri")))
	if err != nil {
		log.Println("Could not connect to DB")
		log.Panicln(err)
	}

	DbClient = newClient
}

func SaveEntry(entry *model.DBEntry) error {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.Replace().SetUpsert(true)

	filter := bson.D{{"_id", entry.ID}}

	_, err := db.Collection("entries").ReplaceOne(context.TODO(), filter, entry, opts)

	return err
}

func GetEntryByHash(hash string) (*model.DBEntry, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection("entries").FindOne(context.TODO(), bson.D{{"hash_value", hash}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := model.DBEntry{}

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}
