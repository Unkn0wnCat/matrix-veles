package web

import (
	"github.com/Unkn0wnCat/matrix-veles/internal/web/api"
	"github.com/Unkn0wnCat/matrix-veles/webui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"time"
)

func StartServer() {
	if viper.GetString("bot.web.secret") == "hunter2" {
		log.Println("Web secret is not set! REFUSING TO START WEB SERVER!")
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/metrics", promhttp.Handler())

	//r.HandleFunc("/", HomeHandler)
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Mount("/api", api.SetupAPI())

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			w.Write([]byte("{\"error_code\": 404, \"error\": \"not_found\"}"))
			return
		}
		return
	})

	ui, err := webui.ServeUI()
	if err == nil {
		r.Mount("/", ui)
	}

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
