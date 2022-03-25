package api

import (
	"context"
	"encoding/json"
	"errors"
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Unkn0wnCat/matrix-veles/graph"
	"github.com/Unkn0wnCat/matrix-veles/graph/generated"
	model2 "github.com/Unkn0wnCat/matrix-veles/graph/model"
	"github.com/Unkn0wnCat/matrix-veles/internal/db/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
	"strings"
)

func SetupAPI() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	m := chiprometheus.NewMiddleware("api")
	router.Use(m)
	router.Use(decodeAuthMiddleware)

	router.NotFound(notFoundHandler)
	router.MethodNotAllowed(methodNotAllowedHandler)

	router.Handle("/", playground.Handler("GraphQL playground", "/api/query"))

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.LoggedIn = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		claimsVal := ctx.Value("claims")
		var claims model.JwtClaims
		if claimsVal != nil {
			claims = claimsVal.(model.JwtClaims)
			if claims.Valid() == nil {
				return next(ctx)
			}
		}

		return nil, nil
	}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model2.UserRole) (res interface{}, err error) {
		user, err := graph.GetUserFromContext(ctx)
		if err != nil {
			if role == model2.UserRoleUnauthenticated {
				return next(ctx)
			}
			return nil, nil
		}

		switch role {
		case model2.UserRoleUser:
			return next(ctx)
		case model2.UserRoleAdmin:
			if user.Admin != nil && *user.Admin {
				return next(ctx)
			}
		case model2.UserRoleUnauthenticated:
			break
		default:
			return nil, errors.New("server error")
		}

		return nil, nil
	}

	c.Directives.Owner = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		user, err := graph.GetUserFromContext(ctx)
		if err != nil {
			return nil, nil
		}

		ctx2 := context.WithValue(ctx, "ownerConstraint", user.ID.Hex())

		return next(ctx2)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	router.Handle("/query", srv)

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

			claims := request.Context().Value("claims").(model.JwtClaims)

			writer.Write([]byte(`hello ` + claims.Username))
		})
	})

	return router
}

func getClaims(request *http.Request) model.JwtClaims {
	claims := request.Context().Value("claims").(model.JwtClaims)

	return claims
}

func decodeAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		tokenSplit := strings.Split(token, " ")

		if token == "" || len(tokenSplit) < 2 {
			next.ServeHTTP(res, req)
			return
		}

		token = tokenSplit[1]

		claims, _, err := parseToken(token)
		if err != nil {
			next.ServeHTTP(res, req)
			return
		}

		ctx := context.WithValue(req.Context(), "claims", *claims)

		req = req.WithContext(ctx)

		next.ServeHTTP(res, req)
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
