package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Unkn0wnCat/matrix-veles/graph/generated"
	"github.com/Unkn0wnCat/matrix-veles/graph/model"
	"github.com/Unkn0wnCat/matrix-veles/internal/config"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	model2 "github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Author is the resolver for the author field.
func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	user, err := db.GetUserByID(*obj.AuthorID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}

		return nil, errors.New("database error")
	}

	return model.MakeUser(user), nil
}

// PartOf is the resolver for the partOf field.
func (r *entryResolver) PartOf(ctx context.Context, obj *model.Entry, first *int, after *string) (*model.ListConnection, error) {
	ids := obj.PartOfIDs

	if len(ids) == 0 {
		return nil, nil
	}

	startIndex := 0

	if after != nil {
		afterInt := new(big.Int)
		afterInt.SetString(*after, 16)

		idInt := new(big.Int)

		set := false

		for i, id := range obj.PartOfIDs {
			idInt.SetString(id.Hex(), 16)

			if idInt.Cmp(afterInt) > 0 {
				startIndex = i
				set = true
				break
			}
		}

		if !set {
			return nil, nil
		}
	}

	if startIndex >= len(ids) {
		return nil, nil
	}

	ids = ids[startIndex:]

	length := 25

	if first != nil {
		length = *first
	}

	cut := false

	if len(ids) > length {
		cut = true
		ids = ids[:length]
	}

	var edges []*model.ListEdge

	for _, id := range ids {
		dbList, err := db.GetListByID(*id)
		if err != nil {
			return nil, err
		}

		edges = append(edges, &model.ListEdge{
			Node:   model.MakeList(dbList),
			Cursor: dbList.ID.Hex(),
		})
	}

	return &model.ListConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: startIndex > 0,
			HasNextPage:     cut,
			StartCursor:     edges[0].Cursor,
			EndCursor:       edges[len(edges)-1].Cursor,
		},
		Edges: edges,
	}, nil
}

// AddedBy is the resolver for the addedBy field.
func (r *entryResolver) AddedBy(ctx context.Context, obj *model.Entry) (*model.User, error) {
	user, err := db.GetUserByID(obj.AddedByID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}

		return nil, errors.New("database error")
	}

	return model.MakeUser(user), nil
}

// Comments is the resolver for the comments field.
func (r *entryResolver) Comments(ctx context.Context, obj *model.Entry, first *int, after *string) (*model.CommentConnection, error) {
	comments := obj.RawComments

	return ResolveComments(comments, first, after)
}

// Creator is the resolver for the creator field.
func (r *listResolver) Creator(ctx context.Context, obj *model.List) (*model.User, error) {
	user, err := db.GetUserByID(obj.CreatorID)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeUser(user), nil
}

// Comments is the resolver for the comments field.
func (r *listResolver) Comments(ctx context.Context, obj *model.List, first *int, after *string) (*model.CommentConnection, error) {
	comments := obj.RawComments

	return ResolveComments(comments, first, after)
}

// Maintainers is the resolver for the maintainers field.
func (r *listResolver) Maintainers(ctx context.Context, obj *model.List, first *int, after *string) (*model.UserConnection, error) {
	ids := obj.MaintainerIDs

	if len(ids) == 0 {
		return nil, nil
	}

	startIndex := 0

	if after != nil {
		afterInt := new(big.Int)
		afterInt.SetString(*after, 16)

		idInt := new(big.Int)

		set := false

		for i, id := range ids {
			idInt.SetString(id.Hex(), 16)

			if idInt.Cmp(afterInt) > 0 {
				startIndex = i
				set = true
				break
			}
		}

		if !set {
			return nil, nil
		}
	}

	if startIndex >= len(ids) {
		return nil, nil
	}

	ids = ids[startIndex:]

	length := 25

	if first != nil {
		length = *first
	}

	cut := false

	if len(ids) > length {
		cut = true
		ids = ids[:length]
	}

	var edges []*model.UserEdge

	for _, id := range ids {
		dbUser, err := db.GetUserByID(*id)
		if err != nil {
			return nil, err
		}

		edges = append(edges, &model.UserEdge{
			Node:   model.MakeUser(dbUser),
			Cursor: dbUser.ID.Hex(),
		})
	}

	return &model.UserConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: startIndex > 0,
			HasNextPage:     cut,
			StartCursor:     edges[0].Cursor,
			EndCursor:       edges[len(edges)-1].Cursor,
		},
		Edges: edges,
	}, nil
}

// Entries is the resolver for the entries field.
func (r *listResolver) Entries(ctx context.Context, obj *model.List, first *int, after *string) (*model.EntryConnection, error) {
	coll := db.Db.Collection(viper.GetString("bot.mongo.collection.entries"))

	dbFilter, _, dbLimit, err := buildDBEntryFilter(first, after, &model.EntryFilter{
		PartOf: &model.IDArrayFilter{
			ContainsAll: []*string{&obj.ID},
		},
	}, nil)
	if err != nil {
		return nil, err
	}

	newLimit := *dbLimit + 1

	findOpts := options.FindOptions{
		Limit: &newLimit,
	}

	res, err := coll.Find(ctx, *dbFilter, &findOpts)
	if err != nil {
		return nil, errors.New("database error")
	}

	var rawEntries []model2.DBEntry

	err = res.All(ctx, &rawEntries)
	if err != nil {
		return nil, errors.New("database error")
	}

	if len(rawEntries) == 0 {
		return nil, nil
	}

	lastEntryI := len(rawEntries) - 1
	if lastEntryI > 0 {
		lastEntryI--
	}

	firstEntry := rawEntries[0]
	lastEntry := rawEntries[lastEntryI]

	isAfter := false

	if after != nil {
		isAfter = true
	}

	hasMore := false
	if int64(len(rawEntries)) > *dbLimit {
		hasMore = true
	}

	var edges []*model.EntryEdge

	for i, rawEntry := range rawEntries {
		if int64(i) == *dbLimit {
			continue
		}

		edges = append(edges, &model.EntryEdge{
			Node:   model.MakeEntry(&rawEntry),
			Cursor: rawEntry.ID.Hex(),
		})
	}

	return &model.EntryConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: isAfter,
			HasNextPage:     hasMore,
			StartCursor:     firstEntry.ID.Hex(),
			EndCursor:       lastEntry.ID.Hex(),
		},
		Edges: edges,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user, err := db.GetUserByUsername(input.Username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("invalid credentials")
		}

		return "", errors.New("database error")
	}

	err = user.CheckPassword(input.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	jwtSigningKey := []byte(viper.GetString("bot.web.secret"))

	claims := model2.JwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365 * 100)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "veles-api",
			Subject:   user.ID.Hex(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", errors.New("unable to create token")
	}

	return ss, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	_, err := db.GetUserByUsername(input.Username)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return "", errors.New("username taken")
	}

	user := model2.DBUser{
		ID:                 primitive.NewObjectID(),
		Username:           input.Username,
		PendingMatrixLinks: []*string{&input.MxID},
		Password:           &input.Password,
	}

	err = user.HashPassword()
	if err != nil {
		return "", errors.New("server error")
	}

	err = db.SaveUser(&user)
	if err != nil {
		return "", errors.New("database error")
	}

	jwtSigningKey := []byte(viper.GetString("bot.web.secret"))

	claims := model2.JwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 365 * 100)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "veles-api",
			Subject:   user.ID.Hex(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", errors.New("unable to create token")
	}

	return ss, nil
}

// AddMxid is the resolver for the addMXID field.
func (r *mutationResolver) AddMxid(ctx context.Context, input model.AddMxid) (*model.User, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	for _, mxid := range append(user.MatrixLinks, user.PendingMatrixLinks...) {
		if strings.EqualFold(*mxid, input.Mxid) {
			return model.MakeUser(user), nil
		}
	}

	user.PendingMatrixLinks = append(user.PendingMatrixLinks, &input.Mxid)

	err = db.SaveUser(user)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeUser(user), nil
}

// RemoveMxid is the resolver for the removeMXID field.
func (r *mutationResolver) RemoveMxid(ctx context.Context, input model.RemoveMxid) (*model.User, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	for i, mxid := range user.MatrixLinks {
		if strings.EqualFold(*mxid, input.Mxid) {
			user.MatrixLinks = append(user.MatrixLinks[:i], user.MatrixLinks[i+1:]...)
		}
	}

	for i, mxid := range user.PendingMatrixLinks {
		if strings.EqualFold(*mxid, input.Mxid) {
			user.PendingMatrixLinks = append(user.PendingMatrixLinks[:i], user.PendingMatrixLinks[i+1:]...)
		}
	}

	err = db.SaveUser(user)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeUser(user), nil
}

// ReconfigureRoom is the resolver for the reconfigureRoom field.
func (r *mutationResolver) ReconfigureRoom(ctx context.Context, input model.RoomConfigUpdate) (*model.Room, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, err
	}

	rConfig, err := config.GetRoomConfigByObjectID(id)
	if err != nil {
		return nil, err
	}

	match := false

	for _, admin := range rConfig.Admins {
		for _, mxid := range user.MatrixLinks {
			if *mxid == admin {
				match = true
			}
		}
	}

	if !match {
		return nil, errors.New("unauthorized")
	}

	if input.Debug != nil {
		rConfig.Debug = *input.Debug
	}

	if input.Deactivate != nil {
		rConfig.Deactivate = *input.Deactivate
	}

	if input.HashChecker != nil {
		if input.HashChecker.HashCheckMode != nil {
			newMode := uint8(0)
			switch *input.HashChecker.HashCheckMode {
			case model.HashCheckerModeNotice:
				newMode = 0
			case model.HashCheckerModeDelete:
				newMode = 1
			case model.HashCheckerModeMute:
				newMode = 2
			case model.HashCheckerModeBan:
				newMode = 3
			default:
				return nil, errors.New("malformed hash check mode")
			}

			rConfig.HashChecker.HashCheckMode = newMode
		}

		if input.HashChecker.ChatNotice != nil {
			rConfig.HashChecker.NoticeToChat = *input.HashChecker.ChatNotice
		}
	}

	if input.AdminPowerLevel != nil {
		if *input.AdminPowerLevel > 100 {
			return nil, errors.New("REFUSING TO SET ADMIN POWER LEVEL > 100")
		}
		rConfig.AdminPowerLevel = *input.AdminPowerLevel
	}

	err = config.SaveRoomConfig(rConfig)

	return model.MakeRoom(rConfig), nil
}

// SubscribeToList is the resolver for the subscribeToList field.
func (r *mutationResolver) SubscribeToList(ctx context.Context, input model.ListSubscriptionUpdate) (*model.Room, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.RoomID)
	if err != nil {
		return nil, err
	}

	rConfig, err := config.GetRoomConfigByObjectID(id)
	if err != nil {
		return nil, err
	}

	match := false

	for _, admin := range rConfig.Admins {
		for _, mxid := range user.MatrixLinks {
			if *mxid == admin {
				match = true
			}
		}
	}

	if !match {
		return nil, errors.New("unauthorized")
	}

	listIdP, err := primitive.ObjectIDFromHex(input.ListID)
	if err != nil {
		return nil, errors.New("unknown list")
	}

	_, err = db.GetListByID(listIdP)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("unknown list")
		}
		return nil, errors.New("database error")
	}

	for _, list := range rConfig.HashChecker.SubscribedLists {
		if list.Hex() == input.ListID {
			return model.MakeRoom(rConfig), nil
		}
	}

	rConfig.HashChecker.SubscribedLists = append(rConfig.HashChecker.SubscribedLists, &listIdP)

	err = config.SaveRoomConfig(rConfig)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeRoom(rConfig), nil
}

// UnsubscribeFromList is the resolver for the unsubscribeFromList field.
func (r *mutationResolver) UnsubscribeFromList(ctx context.Context, input model.ListSubscriptionUpdate) (*model.Room, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.RoomID)
	if err != nil {
		return nil, err
	}

	rConfig, err := config.GetRoomConfigByObjectID(id)
	if err != nil {
		return nil, err
	}

	match := false

	for _, admin := range rConfig.Admins {
		for _, mxid := range user.MatrixLinks {
			if *mxid == admin {
				match = true
			}
		}
	}

	if !match {
		return nil, errors.New("unauthorized")
	}

	listIdP, err := primitive.ObjectIDFromHex(input.ListID)
	if err != nil {
		return nil, errors.New("unknown list")
	}

	_, err = db.GetListByID(listIdP)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("unknown list")
		}
		return nil, errors.New("database error")
	}

	for i, list := range rConfig.HashChecker.SubscribedLists {
		if list.Hex() == input.ListID {
			rConfig.HashChecker.SubscribedLists = append(rConfig.HashChecker.SubscribedLists[:i], rConfig.HashChecker.SubscribedLists[i+1:]...)
			break
		}
	}

	err = config.SaveRoomConfig(rConfig)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeRoom(rConfig), nil
}

// CreateEntry is the resolver for the createEntry field.
func (r *mutationResolver) CreateEntry(ctx context.Context, input model.CreateEntry) (*model.Entry, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	entry, err := db.GetEntryByHash(input.HashValue)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.New("database error")
	}

	if entry == nil {
		entry = &model2.DBEntry{
			ID:        primitive.NewObjectID(),
			Tags:      input.Tags,
			HashValue: input.HashValue,
			Timestamp: time.Now(),
			AddedBy:   &user.ID,
			Comments:  nil,
		}
	}

	if len(input.PartOf) > 0 {
		for _, partOfId := range input.PartOf {
			err = PerformListMaintainerCheck(partOfId, user.ID.Hex())
			if err != nil {
				return nil, errors.New("error adding to lists")
			}

			partOf, _ := primitive.ObjectIDFromHex(partOfId) // This can't fail, it worked in PerformListMaintainerCheck
			entry.AddTo(&partOf)
		}
	}

	if input.Comment != nil {
		entry.Comments = append(entry.Comments, &model2.DBComment{
			Timestamp:   time.Now(),
			CommentedBy: &user.ID,
			Content:     *input.Comment,
		})
	}

	err = db.SaveEntry(entry)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeEntry(entry), nil
}

// CommentEntry is the resolver for the commentEntry field.
func (r *mutationResolver) CommentEntry(ctx context.Context, input model.CommentEntry) (*model.Entry, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.Entry)
	if err != nil {
		return nil, err
	}

	entry, err := db.GetEntryByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}
		return nil, errors.New("database error")
	}

	entry.Comments = append(entry.Comments, &model2.DBComment{
		Timestamp:   time.Now(),
		CommentedBy: &user.ID,
		Content:     input.Comment,
	})

	err = db.SaveEntry(entry)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeEntry(entry), nil
}

// AddToLists is the resolver for the addToLists field.
func (r *mutationResolver) AddToLists(ctx context.Context, input model.AddToLists) (*model.Entry, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.Entry)
	if err != nil {
		return nil, err
	}

	entry, err := db.GetEntryByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}
		return nil, errors.New("database error")
	}

	if len(input.Lists) > 0 {
		for _, partOfId := range input.Lists {
			err = PerformListMaintainerCheck(partOfId, user.ID.Hex())
			if err != nil {
				return nil, errors.New("error adding to lists")
			}

			partOf, _ := primitive.ObjectIDFromHex(partOfId) // This can't fail, it worked in PerformListMaintainerCheck
			entry.AddTo(&partOf)
		}
	}

	err = db.SaveEntry(entry)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeEntry(entry), nil
}

// RemoveFromLists is the resolver for the removeFromLists field.
func (r *mutationResolver) RemoveFromLists(ctx context.Context, input model.RemoveFromLists) (*model.Entry, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.Entry)
	if err != nil {
		return nil, err
	}

	entry, err := db.GetEntryByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}
		return nil, errors.New("database error")
	}

	if len(input.Lists) > 0 {
		for _, partOfId := range input.Lists {
			err = PerformListMaintainerCheck(partOfId, user.ID.Hex())
			if err != nil {
				return nil, errors.New("error adding to lists")
			}

			partOf, _ := primitive.ObjectIDFromHex(partOfId) // This can't fail, it worked in PerformListMaintainerCheck
			entry.RemoveFrom(&partOf)
		}
	}

	err = db.SaveEntry(entry)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeEntry(entry), nil
}

// CreateList is the resolver for the createList field.
func (r *mutationResolver) CreateList(ctx context.Context, input model.CreateList) (*model.List, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.GetListByName(input.Name)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, errors.New("name taken")
	}

	list := model2.DBHashList{
		ID:      primitive.NewObjectID(),
		Name:    input.Name,
		Tags:    input.Tags,
		Creator: user.ID,
	}

	for _, maintainer := range input.Maintainers {
		id, err := primitive.ObjectIDFromHex(maintainer)
		if err != nil {
			return nil, err
		}

		maintainerUser, err := db.GetUserByID(id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("maintainer not found")
			}

			return nil, errors.New("database error")
		}

		list.Maintainers = append(list.Maintainers, &maintainerUser.ID)
	}

	if input.Comment != nil {
		list.Comments = append(list.Comments, &model2.DBComment{
			Timestamp:   time.Now(),
			CommentedBy: &user.ID,
			Content:     *input.Comment,
		})
	}

	err = db.SaveList(&list)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeList(&list), nil
}

// CommentList is the resolver for the commentList field.
func (r *mutationResolver) CommentList(ctx context.Context, input model.CommentList) (*model.List, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(input.List)
	if err != nil {
		return nil, err
	}

	list, err := db.GetListByID(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("not found")
		}
		return nil, errors.New("database error")
	}

	list.Comments = append(list.Comments, &model2.DBComment{
		Timestamp:   time.Now(),
		CommentedBy: &user.ID,
		Content:     input.Comment,
	})

	err = db.SaveList(list)
	if err != nil {
		return nil, errors.New("database error")
	}

	return model.MakeList(list), nil
}

// DeleteList is the resolver for the deleteList field.
func (r *mutationResolver) DeleteList(ctx context.Context, input string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, first *int, after *string, filter *model.UserFilter, sort *model.UserSort) (*model.UserConnection, error) {
	dbFilter, dbSort, dbLimit, err := buildDBUserFilter(first, after, filter, sort)
	if err != nil {
		return nil, err
	}

	coll := db.Db.Collection(viper.GetString("bot.mongo.collection.users"))

	newLimit := *dbLimit + 1

	findOpts := options.FindOptions{
		Limit: &newLimit,
		Sort:  *dbSort,
	}

	res, err := coll.Find(ctx, *dbFilter, &findOpts)
	if err != nil {
		return nil, errors.New("database error")
	}

	var rawUsers []model2.DBUser

	err = res.All(ctx, &rawUsers)
	if err != nil {
		return nil, errors.New("database error")
	}

	if len(rawUsers) == 0 {
		return nil, errors.New("not found")
	}

	lastUserI := len(rawUsers) - 1
	if lastUserI > 0 {
		lastUserI--
	}

	firstUser := rawUsers[0]
	lastUser := rawUsers[lastUserI]

	isAfter := false

	if after != nil {
		isAfter = true
	}

	hasMore := false
	if int64(len(rawUsers)) > *dbLimit {
		hasMore = true
	}

	var edges []*model.UserEdge

	for i, rawUser := range rawUsers {
		if int64(i) == *dbLimit {
			continue
		}

		edges = append(edges, &model.UserEdge{
			Node:   model.MakeUser(&rawUser),
			Cursor: rawUser.ID.Hex(),
		})
	}

	return &model.UserConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: isAfter,
			HasNextPage:     hasMore,
			StartCursor:     firstUser.ID.Hex(),
			EndCursor:       lastUser.ID.Hex(),
		},
		Edges: edges,
	}, nil
}

// Lists is the resolver for the lists field.
func (r *queryResolver) Lists(ctx context.Context, first *int, after *string, filter *model.ListFilter, sort *model.ListSort) (*model.ListConnection, error) {
	dbFilter, dbSort, dbLimit, err := buildDBListFilter(first, after, filter, sort)
	if err != nil {
		return nil, err
	}

	coll := db.Db.Collection(viper.GetString("bot.mongo.collection.lists"))

	newLimit := *dbLimit + 1

	findOpts := options.FindOptions{
		Limit: &newLimit,
		Sort:  *dbSort,
	}

	res, err := coll.Find(ctx, *dbFilter, &findOpts)
	if err != nil {
		return nil, errors.New("database error")
	}

	var rawLists []model2.DBHashList

	err = res.All(ctx, &rawLists)
	if err != nil {
		return nil, errors.New("database error")
	}

	if len(rawLists) == 0 {
		return nil, errors.New("not found")
	}

	lastListI := len(rawLists) - 1
	if lastListI > 0 {
		lastListI--
	}

	firstList := rawLists[0]
	lastList := rawLists[lastListI]

	isAfter := false

	if after != nil {
		isAfter = true
	}

	hasMore := false
	if int64(len(rawLists)) > *dbLimit {
		hasMore = true
	}

	var edges []*model.ListEdge

	for i, rawList := range rawLists {
		if int64(i) == *dbLimit {
			continue
		}

		edges = append(edges, &model.ListEdge{
			Node:   model.MakeList(&rawList),
			Cursor: rawList.ID.Hex(),
		})
	}

	return &model.ListConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: isAfter,
			HasNextPage:     hasMore,
			StartCursor:     firstList.ID.Hex(),
			EndCursor:       lastList.ID.Hex(),
		},
		Edges: edges,
	}, nil
}

// Entries is the resolver for the entries field.
func (r *queryResolver) Entries(ctx context.Context, first *int, after *string, filter *model.EntryFilter, sort *model.EntrySort) (*model.EntryConnection, error) {
	dbFilter, dbSort, dbLimit, err := buildDBEntryFilter(first, after, filter, sort)
	if err != nil {
		return nil, err
	}

	coll := db.Db.Collection(viper.GetString("bot.mongo.collection.entries"))

	newLimit := *dbLimit + 1

	findOpts := options.FindOptions{
		Limit: &newLimit,
		Sort:  *dbSort,
	}

	res, err := coll.Find(ctx, *dbFilter, &findOpts)
	if err != nil {
		return nil, errors.New("database error")
	}

	var rawEntries []model2.DBEntry

	err = res.All(ctx, &rawEntries)
	if err != nil {
		return nil, errors.New("database error")
	}

	if len(rawEntries) == 0 {
		return nil, errors.New("not found")
	}

	lastEntryI := len(rawEntries) - 1
	if lastEntryI > 0 {
		lastEntryI--
	}

	firstEntry := rawEntries[0]
	lastEntry := rawEntries[lastEntryI]

	isAfter := false

	if after != nil {
		isAfter = true
	}

	hasMore := false
	if int64(len(rawEntries)) > *dbLimit {
		hasMore = true
	}

	var edges []*model.EntryEdge

	for i, rawEntry := range rawEntries {
		if int64(i) == *dbLimit {
			continue
		}

		edges = append(edges, &model.EntryEdge{
			Node:   model.MakeEntry(&rawEntry),
			Cursor: rawEntry.ID.Hex(),
		})
	}

	return &model.EntryConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: isAfter,
			HasNextPage:     hasMore,
			StartCursor:     firstEntry.ID.Hex(),
			EndCursor:       lastEntry.ID.Hex(),
		},
		Edges: edges,
	}, nil
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context, first *int, after *string, filter *model.RoomFilter) (*model.RoomConnection, error) {
	user, _ := GetUserFromContext(ctx)

	var userMxids []string

	if user != nil {
		for _, mxid := range user.MatrixLinks {
			userMxids = append(userMxids, *mxid)
		}
	}

	dbFilter, dbSort, dbLimit, err := buildDBRoomFilter(first, after, filter, userMxids)
	if err != nil {
		return nil, err
	}

	coll := db.Db.Collection(viper.GetString("bot.mongo.collection.rooms"))

	newLimit := *dbLimit + 1

	findOpts := options.FindOptions{
		Limit: &newLimit,
		Sort:  *dbSort,
	}

	res, err := coll.Find(ctx, *dbFilter, &findOpts)
	if err != nil {
		return nil, errors.New("database error")
	}

	var rawEntries []config.RoomConfig

	err = res.All(ctx, &rawEntries)
	if err != nil {
		return nil, errors.New("database error")
	}

	if len(rawEntries) == 0 {
		return nil, errors.New("not found")
	}

	lastEntryI := len(rawEntries) - 1
	if lastEntryI > 0 {
		lastEntryI--
	}

	firstEntry := rawEntries[0]
	lastEntry := rawEntries[lastEntryI]

	isAfter := false

	if after != nil {
		isAfter = true
	}

	hasMore := false
	if int64(len(rawEntries)) > *dbLimit {
		hasMore = true
	}

	var edges []*model.RoomEdge

	for i, rawRoom := range rawEntries {
		if int64(i) == *dbLimit {
			continue
		}

		edges = append(edges, &model.RoomEdge{
			Node:   model.MakeRoom(&rawRoom),
			Cursor: rawRoom.ID.Hex(),
		})
	}

	return &model.RoomConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: isAfter,
			HasNextPage:     hasMore,
			StartCursor:     firstEntry.ID.Hex(),
			EndCursor:       lastEntry.ID.Hex(),
		},
		Edges: edges,
	}, nil
}

// Room is the resolver for the room field.
func (r *queryResolver) Room(ctx context.Context, id *string) (*model.Room, error) {
	if id != nil {
		dbId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}

		coll := db.Db.Collection(viper.GetString("bot.mongo.collection.rooms"))

		res := coll.FindOne(ctx, bson.M{"_id": dbId})
		if res.Err() != nil {
			if errors.Is(res.Err(), mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		var room config.RoomConfig

		err = res.Decode(&room)
		if err != nil {
			return nil, errors.New("database error")
		}

		return model.MakeRoom(&room), nil
	}

	return nil, errors.New("not found")
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, username *string) (*model.User, error) {
	if id != nil {
		dbId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}

		rawUser, err := db.GetUserByID(dbId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeUser(rawUser), nil
	}

	if username != nil {
		rawUser, err := db.GetUserByUsername(*username)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeUser(rawUser), nil
	}

	return nil, errors.New("not found")
}

// Entry is the resolver for the entry field.
func (r *queryResolver) Entry(ctx context.Context, id *string, hashValue *string) (*model.Entry, error) {
	if id != nil {
		dbId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}

		rawEntry, err := db.GetEntryByID(dbId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeEntry(rawEntry), nil
	}

	if hashValue != nil {
		rawEntry, err := db.GetEntryByHash(*hashValue)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeEntry(rawEntry), nil
	}

	return nil, errors.New("not found")
}

// List is the resolver for the list field.
func (r *queryResolver) List(ctx context.Context, id *string, name *string) (*model.List, error) {
	if id != nil {
		dbId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}

		rawList, err := db.GetListByID(dbId)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeList(rawList), nil
	}

	if name != nil {
		rawList, err := db.GetListByName(*name)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("not found")
			}

			return nil, errors.New("database error")
		}

		return model.MakeList(rawList), nil
	}

	return nil, errors.New("not found")
}

// Self is the resolver for the self field.
func (r *queryResolver) Self(ctx context.Context) (*model.User, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	return model.MakeUser(user), nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Entry returns generated.EntryResolver implementation.
func (r *Resolver) Entry() generated.EntryResolver { return &entryResolver{r} }

// List returns generated.ListResolver implementation.
func (r *Resolver) List() generated.ListResolver { return &listResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
type entryResolver struct{ *Resolver }
type listResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
