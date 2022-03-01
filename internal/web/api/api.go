package api

import (
	"context"
	"encoding/json"
	"errors"
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func SetupAPI() chi.Router {
	router := chi.NewRouter()

	m := chiprometheus.NewMiddleware("api")
	router.Use(m)

	router.NotFound(notFoundHandler)
	router.MethodNotAllowed(methodNotAllowedHandler)

	//router.NotFoundHandler = NotFoundHandler{}
	//router.MethodNotAllowedHandler = MethodNotAllowedHandler{}

	router.Post("/auth/login", apiHandleAuthLogin)
	router.Post("/auth/register", apiHandleAuthRegister)

	router.Route("/users", func(r chi.Router) {
		r.Use(checkAuthMiddleware)

		r.Get("/", apiHandleAuthUsersList)

		r.Get("/by-name/{name}", apiHandleAuthUserByName)
		//r.Get("/by-mid/{name}", apiHandleAuthUserByMid)
		r.Get("/@self", apiHandleAuthUserSelf)
		r.Get("/{id}", apiHandleAuthUser)
	})

	router.Route("/entries", func(r chi.Router) {
		r.Use(checkAuthMiddleware)

		r.Get("/", apiHandleBotEntriesList)
		r.Post("/", apiHandleBotEntriesPost)

		r.Get("/by-hash/{hash}", apiHandleBotEntryByHash)
		r.Get("/{id}", apiHandleBotEntry)
	})

	router.Route("/lists", func(r chi.Router) {
		r.Use(checkAuthMiddleware)

		r.Get("/", apiHandleBotListsList)
		r.Post("/", apiHandleBotListsPost)

		r.Get("/by-name/{name}", apiHandleBotListByName)
		r.Get("/{id}", apiHandleBotList)
	})

	router.Route("/test", func(r chi.Router) {
		r.Use(checkAuthMiddleware)

		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)

			claims := request.Context().Value("claims").(jwtClaims)

			writer.Write([]byte(`hello ` + claims.Username))
		})
	})

	return router
}

func getClaims(request *http.Request) jwtClaims {
	claims := request.Context().Value("claims").(jwtClaims)

	return claims
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

func notFoundHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	_, _ = res.Write([]byte(`{"error": "not found","error_code":404}`))
}

func methodNotAllowedHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = res.Write([]byte(`{"error": "method not allowed","error_code":405}`))
}
