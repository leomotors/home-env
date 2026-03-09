package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leomotors/home-env/constants"
	"github.com/leomotors/home-env/docs"
	"github.com/leomotors/home-env/middlewares"
	"github.com/leomotors/home-env/routes"
	"github.com/leomotors/home-env/services"
)

//	@title			Home Env API
//	@version		3.1.0
//	@description	Home environment sensor monitoring API
//	@host			localhost:8939
//	@BasePath		/

func healthCheckLoop() {
	for {
		services.HealthCheck()

		time.Sleep(5 * time.Second)
	}
}

func main() {
	secret := services.GetSecret()
	services.InitDB(secret.DATABASE_URL)
	defer services.CloseDB()

	services.RegisterSensor(constants.MainRoomId, "Office Room")
	services.RegisterSensor(constants.LivingRoomId, "Living Room")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		routes.IndexGetHandler.ServeHTTP(w, r)
	})

	mux.Handle("/data", routes.DataGetHandler)
	mux.Handle("/update", middlewares.LocalOnly(routes.UpdatePostHandler))
	mux.Handle("/scalar", routes.ScalarHandler)

	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		w.Write(docs.SwaggerJSON)
	})

	mux.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/yaml")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		w.Write(docs.SwaggerYAML)
	})

	wrappedMux := middlewares.Logger(mux)

	go healthCheckLoop()

	const PORT = 8939
	baseURL := fmt.Sprintf("http://localhost:%d", PORT)
	indexURL := baseURL
	docsURL := fmt.Sprintf("%s/scalar", baseURL)

	const bold = "\033[1m"
	const reset = "\033[0m"

	fmt.Printf("\n🚀 Home Env server is up and running!\n")
	fmt.Printf("- Index page: %s%s%s\n", bold, indexURL, reset)
	fmt.Printf("- Docs page:  %s%s%s\n\n", bold, docsURL, reset)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), wrappedMux))
}
