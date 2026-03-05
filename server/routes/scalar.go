package routes

import (
	"net/http"

	"github.com/leomotors/home-env/public"
)

func scalarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Write(public.ScalarHTML)
}

var ScalarHandler = http.HandlerFunc(scalarHandler)
