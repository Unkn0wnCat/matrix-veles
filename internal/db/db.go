package db

import (
	"context"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	_, err := db.Collection(viper.GetString("bot.mongo.collection.entries")).ReplaceOne(context.TODO(), filter, entry, opts)

	return err
}

func GetEntryByID(id primitive.ObjectID) (*model.DBEntry, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.entries")).FindOne(context.TODO(), bson.D{{"_id", id}})
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

func GetEntries(first int64, cursor *primitive.ObjectID) ([]*model.DBEntry, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.FindOptions{
		Limit: &first,
	}

	filter := bson.M{}

	if cursor != nil {
		filter = bson.M{
			"_id": bson.M{
				"$gt": *cursor,
			},
		}
		log.Println(filter)
	}

	res, err := db.Collection(viper.GetString("bot.mongo.collection.entries")).Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, res.Err()
	}

	var object []*model.DBEntry

	err = res.All(context.TODO(), &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func GetEntryByHash(hash string) (*model.DBEntry, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.entries")).FindOne(context.TODO(), bson.D{{"hash_value", hash}})
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

func SaveList(list *model.DBHashList) error {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.Replace().SetUpsert(true)

	filter := bson.D{{"_id", list.ID}}

	_, err := db.Collection(viper.GetString("bot.mongo.collection.lists")).ReplaceOne(context.TODO(), filter, list, opts)

	return err
}

func GetListByID(id primitive.ObjectID) (*model.DBHashList, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.lists")).FindOne(context.TODO(), bson.D{{"_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := model.DBHashList{}

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func GetListByName(name string) (*model.DBHashList, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.lists")).FindOne(context.TODO(), bson.D{{"name", name}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := model.DBHashList{}

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func GetLists(first int64, cursor *primitive.ObjectID) ([]*model.DBHashList, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.FindOptions{
		Limit: &first,
	}

	filter := bson.M{}

	if cursor != nil {
		filter = bson.M{
			"_id": bson.M{
				"$gt": *cursor,
			},
		}
		log.Println(filter)
	}

	res, err := db.Collection(viper.GetString("bot.mongo.collection.lists")).Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, res.Err()
	}

	var object []*model.DBHashList

	err = res.All(context.TODO(), &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func SaveUser(user *model.DBUser) error {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.Replace().SetUpsert(true)

	filter := bson.D{{"_id", user.ID}}

	_, err := db.Collection(viper.GetString("bot.mongo.collection.users")).ReplaceOne(context.TODO(), filter, user, opts)

	return err
}

func GetUserByID(id primitive.ObjectID) (*model.DBUser, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.users")).FindOne(context.TODO(), bson.D{{"_id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := model.DBUser{}

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func GetUserByUsername(username string) (*model.DBUser, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	res := db.Collection(viper.GetString("bot.mongo.collection.users")).FindOne(context.TODO(), bson.D{{"username", username}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	object := model.DBUser{}

	err := res.Decode(&object)
	if err != nil {
		return nil, err
	}

	return &object, nil
}

func GetUsers(first int64, cursor *primitive.ObjectID) ([]*model.DBUser, error) {
	db := DbClient.Database(viper.GetString("bot.mongo.database"))

	opts := options.FindOptions{
		Limit: &first,
	}

	filter := bson.M{}

	if cursor != nil {
		filter = bson.M{
			"_id": bson.M{
				"$gt": *cursor,
			},
		}
		log.Println(filter)
	}

	res, err := db.Collection(viper.GetString("bot.mongo.collection.users")).Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, res.Err()
	}

	var object []*model.DBUser

	err = res.All(context.TODO(), &object)
	if err != nil {
		return nil, err
	}

	return object, nil
}
