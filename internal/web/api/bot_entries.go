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

type apiEntryPostBody struct {
	Hash    string                `json:"hash"`
	Tags    []string              `json:"tags"`
	PartOf  []*primitive.ObjectID `json:"part_of"`
	FileURL string                `json:"file_url"`
	Comment *string               `json:"comment"`
}

func apiHandleBotEntriesPost(res http.ResponseWriter, req *http.Request) {
	var body apiEntryPostBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("malformed body"))
		return
	}

	existingEntry, err := db.GetEntryByHash(body.Hash)
	if err == nil {
		writeJSONError(res, http.StatusConflict, fmt.Errorf("hash already in database: %s", existingEntry.ID))
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

	for i, partOf := range body.PartOf {
		list, err := db.GetListByID(*partOf)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				writeJSONError(res, http.StatusNotFound, fmt.Errorf("invalid partof value at index %d - not found", i))
				return
			}
			// TODO: LOG THIS ERROR
			writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
			return
		}

		authorized := false

		for _, maintainer := range list.Maintainers {
			log.Println(maintainer, userId)
			if maintainer.Hex() == userId.Hex() {
				authorized = true
				break
			}
		}

		if !authorized {
			writeJSONError(res, http.StatusUnauthorized, fmt.Errorf("invalid partof value at index %d - not authorized", i))
			return
		}
	}

	newEntry := model.DBEntry{
		ID:        primitive.NewObjectID(),
		Tags:      body.Tags,
		PartOf:    body.PartOf,
		HashValue: body.Hash,
		FileURL:   body.FileURL,
		Timestamp: time.Now(),
		AddedBy:   &userId,
		Comments:  nil,
	}

	if body.Comment != nil && *body.Comment != "" {
		newEntry.Comments = append(newEntry.Comments, &model.DBComment{
			CommentedBy: &userId,
			Content:     *body.Comment,
			Timestamp:   time.Now(),
		})
	}

	err = db.SaveEntry(&newEntry)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	encoded, err := json.Marshal(newEntry)
	if err != nil {
		// TODO: LOG THIS ERROR
		writeJSONError(res, http.StatusInternalServerError, errors.New("could not marshal data"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(encoded)
}

func apiHandleBotEntriesList(res http.ResponseWriter, req *http.Request) {
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

	entries, err := db.GetEntries(first, cursor)
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

func apiHandleBotEntry(res http.ResponseWriter, req *http.Request) {
	requestedId := chi.URLParam(req, "id")
	objectId, err := primitive.ObjectIDFromHex(requestedId)
	if err != nil {
		writeJSONError(res, http.StatusNotFound, errors.New("malformed id"))
		return
	}

	entry, err := db.GetEntryByID(objectId)
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

func apiHandleBotEntryByHash(res http.ResponseWriter, req *http.Request) {
	requestedHash := chi.URLParam(req, "hash")

	entry, err := db.GetEntryByHash(requestedHash)
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
