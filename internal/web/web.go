package web

import (
	"github.com/Unkn0wnCat/matrix-veles/internal/web/api"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	apiRouter := r.PathPrefix("/api").Subrouter()
	api.SetupAPI(apiRouter)

	srv := &http.Server{
		Handler:      r,
		Addr:         viper.GetString("bot.web.listen"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Now serving web-interface on http://%s", viper.GetString("bot.web.listen"))

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("hello world"))
}
