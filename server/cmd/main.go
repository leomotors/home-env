package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leomotors/home-env/constants"
	"github.com/leomotors/home-env/middlewares"
	"github.com/leomotors/home-env/routes"
	"github.com/leomotors/home-env/services"
)

func healthCheckLoop() {
	for {
		services.HealthCheck()

		time.Sleep(5 * time.Second)
	}
}

func main() {
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
	mux.Handle("/metrics", routes.MetricsHandler)
	mux.Handle("/update", routes.UpdatePostHandler)

	wrappedMux := middlewares.Logger(mux)

	go healthCheckLoop()

	const PORT = 8939
	fmt.Printf("Listening on port %d...\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), wrappedMux))
}
