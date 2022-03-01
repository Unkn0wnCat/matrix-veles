package api

import (
	"encoding/json"
	"errors"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/url"
	"strconv"
)

func apiHandleAuthUsersList(res http.ResponseWriter, req *http.Request) {
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

	entries, err := db.GetUsers(first, cursor)
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

func apiHandleAuthUser(res http.ResponseWriter, req *http.Request) {
	requestedId := chi.URLParam(req, "id")
	objectId, err := primitive.ObjectIDFromHex(requestedId)
	if err != nil {
		writeJSONError(res, http.StatusNotFound, errors.New("malformed id"))
		return
	}

	entry, err := db.GetUserByID(objectId)
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

func apiHandleAuthUserSelf(res http.ResponseWriter, req *http.Request) {
	claims := getClaims(req)
	objectId, err := primitive.ObjectIDFromHex(claims.Subject)
	if err != nil {
		writeJSONError(res, http.StatusNotFound, errors.New("malformed id"))
		return
	}

	entry, err := db.GetUserByID(objectId)
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

func apiHandleAuthUserByName(res http.ResponseWriter, req *http.Request) {
	requestedName := chi.URLParam(req, "name")

	entry, err := db.GetUserByUsername(requestedName)
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
