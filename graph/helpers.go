package graph

import (
	"context"
	"errors"
	"github.com/Unkn0wnCat/matrix-veles/graph/model"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	model2 "github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func PerformListMaintainerCheck(listIdHex string, userIdHex string) error {
	id, err := primitive.ObjectIDFromHex(listIdHex)
	if err != nil {
		return err
	}

	list, err := db.GetListByID(id)
	if err != nil {
		return err
	}

	if list.Creator.Hex() == userIdHex {
		return nil
	}

	for _, maintainerId := range list.Maintainers {
		if maintainerId.Hex() == userIdHex {
			return nil
		}
	}

	return errors.New("unauthorized")
}

func GetUserFromContext(ctx context.Context) (*model2.DBUser, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := db.GetUserByID(*userID)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	return user, nil
}

func GetUserIDFromContext(ctx context.Context) (*primitive.ObjectID, error) {
	claimsVal := ctx.Value("claims")
	var claims model2.JwtClaims
	if claimsVal != nil {
		claims = claimsVal.(model2.JwtClaims)
		if claims.Valid() == nil {
			sub := claims.Subject

			id, err := primitive.ObjectIDFromHex(sub)
			if err != nil {
				return nil, err
			}

			return &id, nil
		}
	}

	return nil, errors.New("no user")
}

func CheckOwnerConstraint(ctx context.Context, realOwnerIDHex string) error {
	constraint, ok := ctx.Value("ownerConstraint").(string)

	if !ok {
		return nil // This means no ownerConstraint is set
	}

	if constraint != realOwnerIDHex {
		return errors.New("owner constraint violation")
	}

	return nil
}

func buildStringFilter(filter *model.StringFilter) bson.M {
	compiledFilter := bson.M{}

	if filter != nil {
		if filter.Eq != nil {
			compiledFilter["$eq"] = *filter.Eq
		}
		if filter.Neq != nil {
			compiledFilter["$ne"] = *filter.Eq
		}
		if filter.Regex != nil {
			compiledFilter["$regex"] = *filter.Regex
		}
	}

	return compiledFilter
}

func buildTimestampFilter(filter *model.TimestampFilter) (*bson.M, error) {
	compiledFilter := bson.M{}

	if filter != nil {
		if filter.After != nil {
			compiledFilter["$gt"] = *filter.After
		}
		if filter.Before != nil {
			compiledFilter["$lt"] = *filter.Before
		}
	}

	return &compiledFilter, nil
}

func buildIntFilter(filter *model.IntFilter) bson.M {
	compiledFilter := bson.M{}

	if filter != nil {
		if filter.Eq != nil {
			compiledFilter["$eq"] = *filter.Eq
		}
		if filter.Neq != nil {
			compiledFilter["$ne"] = *filter.Neq
		}
		if filter.Gt != nil {
			compiledFilter["$gt"] = *filter.Gt
		}
		if filter.Lt != nil {
			compiledFilter["$lt"] = *filter.Lt
		}
	}

	return compiledFilter
}

func buildStringArrayFilter(filter *model.StringArrayFilter) bson.M {
	compiledFilter := bson.M{}

	if filter != nil {
		if filter.ElemMatch != nil {
			compiledFilter["$elemMatch"] = buildStringFilter(filter.ElemMatch)
		}
		if filter.ContainsAll != nil {
			compiledFilter["$all"] = filter.ContainsAll
		}
		if filter.Length != nil {
			compiledFilter["$size"] = *filter.Length
		}
	}

	return compiledFilter
}

func idArrayToPrimitiveID(ids []*string) ([]primitive.ObjectID, error) {
	var pIds []primitive.ObjectID

	for _, id := range ids {
		pId, err := primitive.ObjectIDFromHex(*id)
		if err != nil {
			return nil, err
		}

		pIds = append(pIds, pId)
	}

	return pIds, nil
}

func buildIDArrayFilter(filter *model.IDArrayFilter) (bson.M, error) {
	compiledFilter := bson.M{}
	var err error

	if filter != nil {
		if filter.ContainsAll != nil {
			compiledFilter["$all"], err = idArrayToPrimitiveID(filter.ContainsAll)
			if err != nil {
				return nil, err
			}
		}
		if filter.Length != nil {
			compiledFilter["$size"] = *filter.Length
		}
	}

	return compiledFilter, nil
}

func buildSortRule(sort *model.SortRule) interface{} {
	if sort.Direction == "ASC" {
		return 1
	}

	return -1
}

func buildDBUserFilter(first *int, after *string, filter *model.UserFilter, sort *model.UserSort) (*bson.M, *bson.M, *int64, error) {
	compiledFilter := bson.M{}
	compiledSort := bson.M{}

	var filterBson *bson.M
	var cursorBson *bson.M
	limit := 25

	if sort != nil {
		if sort.ID != nil {
			compiledSort["_id"] = buildSortRule(sort.ID)
		}
		if sort.Username != nil {
			compiledSort["username"] = buildSortRule(sort.Username)
		}
		if sort.Admin != nil {
			compiledSort["admin"] = buildSortRule(sort.Admin)
		}
	}

	if first != nil {
		limit = *first
	}

	if after != nil {
		cursorBsonW := bson.M{}

		afterID, err := primitive.ObjectIDFromHex(*after)
		if err != nil {
			return nil, nil, nil, err
		}

		cursorBsonW["_id"] = bson.M{"$gt": afterID}

		cursorBson = &cursorBsonW
	}

	if filter != nil {
		filterBsonW := bson.M{}

		if filter.ID != nil {
			filterBsonW["_id"] = *filter.ID
		}

		if filter.Username != nil {
			filterBsonW["username"] = buildStringFilter(filter.Username)
		}

		if filter.Admin != nil {
			filterBsonW["admin"] = *filter.Admin
		}

		if filter.MatrixLinks != nil {
			filterBsonW["matrix_links"] = buildStringArrayFilter(filter.MatrixLinks)
		}

		if filter.PendingMatrixLinks != nil {
			filterBsonW["pending_matrix_links"] = buildStringArrayFilter(filter.PendingMatrixLinks)
		}

		filterBson = &filterBsonW
	}

	if filterBson != nil && cursorBson != nil {
		compiledFilter["$and"] = bson.A{*cursorBson, *filterBson}
	}

	if filterBson == nil && cursorBson != nil {
		compiledFilter = *cursorBson
	}

	if filterBson != nil && cursorBson == nil {
		compiledFilter = *filterBson
	}

	convLimit := int64(limit)

	return &compiledFilter, &compiledSort, &convLimit, nil
}

func buildDBListFilter(first *int, after *string, filter *model.ListFilter, sort *model.ListSort) (*bson.M, *bson.M, *int64, error) {
	compiledFilter := bson.M{}
	compiledSort := bson.M{}

	var filterBson *bson.M
	var cursorBson *bson.M
	limit := 25

	var err error

	if sort != nil {
		if sort.ID != nil {
			compiledSort["_id"] = buildSortRule(sort.ID)
		}
		if sort.Name != nil {
			compiledSort["name"] = buildSortRule(sort.Name)
		}
	}

	if first != nil {
		limit = *first
	}

	if after != nil {
		cursorBsonW := bson.M{}

		afterID, err := primitive.ObjectIDFromHex(*after)
		if err != nil {
			return nil, nil, nil, err
		}

		cursorBsonW["_id"] = bson.M{"$gt": afterID}

		cursorBson = &cursorBsonW
	}

	if filter != nil {
		filterBsonW := bson.M{}

		if filter.ID != nil {
			filterBsonW["_id"] = *filter.ID
		}

		if filter.Name != nil {
			filterBsonW["name"] = buildStringFilter(filter.Name)
		}

		if filter.Tags != nil {
			filterBsonW["tags"] = buildStringArrayFilter(filter.Tags)
		}

		if filter.Maintainers != nil {
			filterBsonW["maintainers"], err = buildIDArrayFilter(filter.Maintainers)
			if err != nil {
				return nil, nil, nil, err
			}
		}

		filterBson = &filterBsonW
	}

	if filterBson != nil && cursorBson != nil {
		compiledFilter["$and"] = bson.A{*cursorBson, *filterBson}
	}

	if filterBson == nil && cursorBson != nil {
		compiledFilter = *cursorBson
	}

	if filterBson != nil && cursorBson == nil {
		compiledFilter = *filterBson
	}

	convLimit := int64(limit)

	return &compiledFilter, &compiledSort, &convLimit, nil
}

func buildDBEntryFilter(first *int, after *string, filter *model.EntryFilter, sort *model.EntrySort) (*bson.M, *bson.M, *int64, error) {
	compiledFilter := bson.M{}
	compiledSort := bson.M{}

	var filterBson *bson.M
	var cursorBson *bson.M
	limit := 25

	var err error

	if sort != nil {
		if sort.ID != nil {
			compiledSort["_id"] = buildSortRule(sort.ID)
		}
		if sort.Timestamp != nil {
			compiledSort["timestamp"] = buildSortRule(sort.Timestamp)
		}
		if sort.AddedBy != nil {
			compiledSort["added_by"] = buildSortRule(sort.AddedBy)
		}
		if sort.HashValue != nil {
			compiledSort["hash_value"] = buildSortRule(sort.HashValue)
		}
	}

	if first != nil {
		limit = *first
	}

	if after != nil {
		cursorBsonW := bson.M{}

		afterID, err := primitive.ObjectIDFromHex(*after)
		if err != nil {
			return nil, nil, nil, err
		}

		cursorBsonW["_id"] = bson.M{"$gt": afterID}

		cursorBson = &cursorBsonW
	}

	if filter != nil {
		filterBsonW := bson.M{}

		if filter.ID != nil {
			filterBsonW["_id"] = *filter.ID
		}

		if filter.HashValue != nil {
			filterBsonW["hash_value"] = buildStringFilter(filter.HashValue)
		}

		if filter.Tags != nil {
			filterBsonW["tags"] = buildStringArrayFilter(filter.Tags)
		}

		if filter.AddedBy != nil {
			dbId, err := primitive.ObjectIDFromHex(*filter.AddedBy)
			if err != nil {
				return nil, nil, nil, err
			}

			filterBsonW["added_by"] = dbId
		}

		if filter.Timestamp != nil {
			filterBsonW["timestamp"], err = buildTimestampFilter(filter.Timestamp)
			if err != nil {
				return nil, nil, nil, err
			}
		}

		/*if filter.FileURL != nil {
			filterBsonW["file_url"] = buildStringFilter(filter.FileURL)
		}*/

		if filter.PartOf != nil {
			filterBsonW["part_of"], err = buildIDArrayFilter(filter.PartOf)
			if err != nil {
				return nil, nil, nil, err
			}
		}

		filterBson = &filterBsonW
	}

	if filterBson != nil && cursorBson != nil {
		compiledFilter["$and"] = bson.A{*cursorBson, *filterBson}
	}

	if filterBson == nil && cursorBson != nil {
		compiledFilter = *cursorBson
	}

	if filterBson != nil && cursorBson == nil {
		compiledFilter = *filterBson
	}

	convLimit := int64(limit)

	return &compiledFilter, &compiledSort, &convLimit, nil
}

func ResolveComments(comments []*model2.DBComment, first *int, after *string) (*model.CommentConnection, error) {
	if len(comments) == 0 {
		return nil, nil
	}

	startIndex := 0

	if after != nil {
		afterTs, err := time.Parse(time.RFC3339Nano, *after)
		if err != nil {
			return nil, err
		}

		set := false

		for i, comment := range comments {
			if afterTs.Before(comment.Timestamp) {
				startIndex = i
				set = true
				break
			}
		}

		if !set {
			return nil, nil
		}
	}

	if startIndex >= len(comments) {
		return nil, nil
	}

	comments = comments[startIndex:]

	length := 25

	if first != nil {
		length = *first
	}

	cut := false

	if len(comments) > length {
		cut = true
		comments = comments[:length]
	}

	var edges []*model.CommentEdge

	for _, comment := range comments {
		edges = append(edges, &model.CommentEdge{
			Node:   model.MakeComment(comment),
			Cursor: comment.Timestamp.Format(time.RFC3339Nano),
		})
	}
	log.Println(edges)

	return &model.CommentConnection{
		PageInfo: &model.PageInfo{
			HasPreviousPage: startIndex > 0,
			HasNextPage:     cut,
			StartCursor:     edges[0].Cursor,
			EndCursor:       edges[len(edges)-1].Cursor,
		},
		Edges: nil,
	}, nil
}

func buildDBRoomFilter(first *int, after *string, filter *model.RoomFilter /*, sort *model.UserSort*/, userId primitive.ObjectID) (*bson.M, *bson.M, *int64, error) {
	compiledFilter := bson.M{}
	compiledSort := bson.M{}

	var filterBson *bson.M
	var cursorBson *bson.M
	limit := 25

	/*if sort != nil {
		if sort.ID != nil {
			compiledSort["_id"] = buildSortRule(sort.ID)
		}
		if sort.Username != nil {
			compiledSort["username"] = buildSortRule(sort.Username)
		}
		if sort.Admin != nil {
			compiledSort["admin"] = buildSortRule(sort.Admin)
		}
	}*/

	if first != nil {
		limit = *first
	}

	if after != nil {
		cursorBsonW := bson.M{}

		afterID, err := primitive.ObjectIDFromHex(*after)
		if err != nil {
			return nil, nil, nil, err
		}

		cursorBsonW["_id"] = bson.M{"$gt": afterID}

		cursorBson = &cursorBsonW
	}

	if filter != nil {
		filterBsonW := bson.M{}

		if filter.ID != nil {
			filterBsonW["_id"] = *filter.ID
		}

		if filter.Debug != nil {
			filterBsonW["debug"] = *filter.Debug
		}

		if filter.Active != nil {
			filterBsonW["active"] = *filter.Active
		}

		if filter.CanEdit != nil && *filter.CanEdit == true {
			filterBsonW["admins"] = userId
		}

		filterBson = &filterBsonW
	}

	if filterBson != nil && cursorBson != nil {
		compiledFilter["$and"] = bson.A{*cursorBson, *filterBson}
	}

	if filterBson == nil && cursorBson != nil {
		compiledFilter = *cursorBson
	}

	if filterBson != nil && cursorBson == nil {
		compiledFilter = *filterBson
	}

	convLimit := int64(limit)

	return &compiledFilter, &compiledSort, &convLimit, nil
}
