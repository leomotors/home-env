package routes

import (
	"net/http"

	"github.com/leomotors/home-env/public"
)

func indexRenderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(public.IndexHTML))
}

var IndexGetHandler = http.HandlerFunc(indexRenderHandler)
