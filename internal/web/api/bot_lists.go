package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type apiListPostBody struct {
	Name        string
	Tags        []string
	Comment     *string
	Maintainers []*primitive.ObjectID
}

func apiHandleBotListsPost(res http.ResponseWriter, req *http.Request) {
	var body apiListPostBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("malformed body"))
		return
	}

	existingEntry, err := db.GetListByName(body.Name)
	if err == nil {
		writeJSONError(res, http.StatusConflict, fmt.Errorf("name taken: %s", existingEntry.ID))
		return
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	user := getClaims(req)
	userId, err := primitive.ObjectIDFromHex(user.Subject)
	if err != nil {
		// TODO: LOG THIS ERROR
		log.Println(userId)
		writeJSONError(res, http.StatusInternalServerError, errors.New("internal corruption 0x01"))
		return
	}

	newList := model.DBHashList{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		Tags:        body.Tags,
		Comments:    nil,
		Maintainers: body.Maintainers,
	}

	listedSelf := false

	for _, maintainer := range newList.Maintainers {
		if maintainer.Hex() == userId.Hex() {
			listedSelf = true
			break
		}
	}

	if !listedSelf {
		newList.Maintainers = append(newList.Maintainers, &userId)
	}

	if body.Comment != nil && *body.Comment != "" {
		newList.Comments = append(newList.Comments, &model.DBComment{
			CommentedBy: &userId,
			Content:     *body.Comment,
			Timestamp:   time.Now(),
		})
	}

	err = db.SaveList(&newList)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	encoded, err := json.Marshal(newList)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("could not marshal data"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(encoded)
}

func apiHandleBotListsList(res http.ResponseWriter, req *http.Request) {
	requestUri, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("unable to parse uri"))
		return
	}

	first := int64(50)
	var cursor *primitive.ObjectID

	if requestUri.Query().Has("first") {
		first2, err := strconv.Atoi(requestUri.Query().Get("first"))
		if err != nil {
			writeJSONError(res, http.StatusBadRequest, errors.New("malformed query"))
			return
		}
		first = int64(first2)
	}

	if requestUri.Query().Has("cursor") {
		cursor2, err := primitive.ObjectIDFromHex(requestUri.Query().Get("cursor"))
		if err != nil {
			writeJSONError(res, http.StatusBadRequest, errors.New("malformed query"))
			return
		}
		cursor = &cursor2
	}

	entries, err := db.GetLists(first, cursor)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			// TODO: LOG THIS ERROR
			writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("[]"))
		return
	}

	encoded, err := json.Marshal(entries)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("could not marshal data"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(encoded)
}

func apiHandleBotList(res http.ResponseWriter, req *http.Request) {
	requestedId := chi.URLParam(req, "id")
	objectId, err := primitive.ObjectIDFromHex(requestedId)
	if err != nil {
		writeJSONError(res, http.StatusNotFound, errors.New("malformed id"))
		return
	}

	entry, err := db.GetListByID(objectId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			writeJSONError(res, http.StatusNotFound, errors.New("not found"))
			return
		}
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	encoded, err := json.Marshal(entry)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("could not marshal data"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(encoded)
}

func apiHandleBotListByName(res http.ResponseWriter, req *http.Request) {
	requestedName := chi.URLParam(req, "name")

	entry, err := db.GetListByName(requestedName)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			writeJSONError(res, http.StatusNotFound, errors.New("not found"))
			return
		}
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	encoded, err := json.Marshal(entry)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("could not marshal data"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(encoded)
}
