package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func SetupAPI(router *mux.Router) {
	router.NotFoundHandler = NotFoundHandler{}
	router.MethodNotAllowedHandler = MethodNotAllowedHandler{}

	router.Path("/auth/login").Methods("POST").HandlerFunc(apiHandleAuthLogin)
	router.Path("/auth/register").Methods("POST").HandlerFunc(apiHandleAuthRegister)

	bot := router.PathPrefix("/bot").Subrouter()
	bot.Use(checkAuthMiddleware)

	bot.Path("/test").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)

		claims := request.Context().Value("claims").(jwtClaims)

		writer.Write([]byte(`hello ` + claims.Username))
	})
}

func checkAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		tokenSplit := strings.Split(token, " ")

		if token == "" || len(tokenSplit) < 2 {
			writeJSONError(res, http.StatusUnauthorized, errors.New("bearer token required"))
			return
		}

		token = tokenSplit[1]

		claims, _, err := parseToken(token)
		if err != nil {
			writeJSONError(res, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx := context.WithValue(req.Context(), "claims", *claims)

		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
	})
}

func writeJSONError(res http.ResponseWriter, statusCode int, err error) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	enc, _ := json.Marshal(struct {
		Error     string `json:"error"`
		ErrorCode int    `json:"error_code"`
	}{
		Error:     err.Error(),
		ErrorCode: statusCode,
	})

	_, _ = res.Write(enc)
}

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	_, _ = res.Write([]byte(`{"error": "not_found","error_code":404}`))
}

type MethodNotAllowedHandler struct{}

func (MethodNotAllowedHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = res.Write([]byte(`{"error": "method_not_allowed","error_code":405}`))
}
