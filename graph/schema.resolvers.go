package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Unkn0wnCat/matrix-veles/graph/generated"
	"github.com/Unkn0wnCat/matrix-veles/graph/model"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	model2 "github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func (r *entryResolver) Comments(ctx context.Context, obj *model.Entry, first *int, after *string) (*model.CommentConnection, error) {
	comments := obj.RawComments

	return ResolveComments(comments, first, after)
}

func (r *listResolver) Creator(ctx context.Context, obj *model.List) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *listResolver) Comments(ctx context.Context, obj *model.List, first *int, after *string) (*model.CommentConnection, error) {
	comments := obj.RawComments

	return ResolveComments(comments, first, after)
}

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

func (r *mutationResolver) Register(ctx context.Context, input model.Register) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMxid(ctx context.Context, input model.AddMxid) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveMxid(ctx context.Context, input model.RemoveMxid) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEntry(ctx context.Context, input model.CreateEntry) (*model.Entry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CommentEntry(ctx context.Context, input model.CommentEntry) (*model.Entry, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEntry(ctx context.Context, input string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateList(ctx context.Context, input model.CreateList) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CommentList(ctx context.Context, input model.CommentList) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddToList(ctx context.Context, input model.AddToList) (*model.List, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteList(ctx context.Context, input string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

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
