package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/leomotors/home-env/middlewares"
	"github.com/leomotors/home-env/routes"
	"github.com/leomotors/home-env/services"
)

func loop() {
	for {
		services.HealthCheck()

		time.Sleep(5 * time.Second)
	}
}

func main() {
	secret := services.GetSecret()

	services.RegisterSensor("main_room", "Office Room")
	services.RegisterSensor("living_room", "Living Room")

	mux := http.NewServeMux()
	mux.Handle("/metrics", routes.MetricsGetHandler)
	mux.Handle("/update", routes.UpdatePostHandler)

	wrappedMux := middlewares.Logger(mux)

	go loop()

	fmt.Printf("Listening on port %d...\n", secret.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", secret.PORT), wrappedMux))
}
