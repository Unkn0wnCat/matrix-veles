//go:build !withUI
// +build !withUI

package webui

//go:generate yarn
//go:generate yarn build

import (
	"net/http"
)

func ServeUI() (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte("{\"error_code\": 404, \"error\": \"not_found\", \"note\": \"WebUI not included in build - visit https://veles.1in1.net/docs/tutorial-basics/install\"}"))
		return
	}), nil
}
