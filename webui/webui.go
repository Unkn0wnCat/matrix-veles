//go:build withUI
// +build withUI

package webui

//go:generate yarn
//go:generate yarn build

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"
	"strings"
)

//go:embed build/*
var content embed.FS

func ServeUI() (http.Handler, error) {
	fSys, err := fs.Sub(content, "build")
	if err != nil {
		return nil, err
	}

	staticServer := http.FileServer(http.FS(fSys))

	myServeIndex, err := serveIndex()
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fSys.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err != nil {
			log.Println(err)
			myServeIndex.ServeHTTP(w, r)
			return
		}
		log.Println("serving static")
		staticServer.ServeHTTP(w, r)
	}), nil
}

func serveIndex() (http.HandlerFunc, error) {
	indexFile, err := content.ReadFile("build/index.html")
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		w.WriteHeader(200)
		w.Write(indexFile)
	}, nil
}
