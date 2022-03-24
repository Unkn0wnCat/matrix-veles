package api

import (
	"encoding/json"
	"errors"
	"github.com/Unkn0wnCat/matrix-veles/internal/db"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type apiAuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func parseToken(tokenString string) (*model.JwtClaims, *jwt.Token, error) {
	claims := model.JwtClaims{}
	jwtSigningKey := []byte(viper.GetString("bot.web.secret"))

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSigningKey, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return &claims, token, nil
}

func apiHandleAuthLogin(res http.ResponseWriter, req *http.Request) {
	body := req.Body

	bodyContent := apiAuthRequestBody{}

	err := json.NewDecoder(body).Decode(&bodyContent)
	if err != nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("malformed body"))
		return
	}

	user, err := db.GetUserByUsername(bodyContent.Username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			writeJSONError(res, http.StatusUnauthorized, errors.New("invalid credentials"))
			return
		}

		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	err = user.CheckPassword(bodyContent.Password)
	if err != nil {
		writeJSONError(res, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	jwtSigningKey := []byte(viper.GetString("bot.web.secret"))

	claims := model.JwtClaims{
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
		writeJSONError(res, http.StatusInternalServerError, errors.New("unable to create token"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	enc, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		Token: ss,
	})

	_, _ = res.Write(enc)
}

type apiAuthRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func apiHandleAuthRegister(res http.ResponseWriter, req *http.Request) {
	body := req.Body

	bodyContent := apiAuthRegisterBody{}

	err := json.NewDecoder(body).Decode(&bodyContent)
	if err != nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("malformed body"))
		return
	}

	_, err = db.GetUserByUsername(bodyContent.Username)
	if err == nil {
		writeJSONError(res, http.StatusBadRequest, errors.New("username taken"))
		return
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	user := model.DBUser{
		ID:       primitive.NewObjectID(),
		Username: bodyContent.Username,
		Password: &bodyContent.Password,
	}

	err = user.HashPassword()
	if err != nil {
		writeJSONError(res, http.StatusInternalServerError, errors.New("unable to hash password"))
		return
	}

	err = db.SaveUser(&user)
	if err != nil {
		writeJSONError(res, http.StatusInternalServerError, errors.New("database error"))
		return
	}

	jwtSigningKey := viper.GetString("bot.web.secret")

	claims := model.JwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "veles-api",
			Subject:   user.ID.Hex(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtSigningKey)
	if err != nil {
		writeJSONError(res, http.StatusInternalServerError, errors.New("unable to create token"))
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	enc, err := json.Marshal(struct {
		Token string `json:"token"`
	}{
		Token: ss,
	})

	_, _ = res.Write(enc)
}
