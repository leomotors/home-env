package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leomotors/home-env/src/services"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check authorization header
	secret := services.GetSecret()
	expectedPassword := secret.PASSWORD
	providedPassword := r.Header.Get("Authorization")

	if providedPassword != expectedPassword {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	temp, ok := data["temperature"].(float64)
	if !ok {
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	hum, ok := data["humidity"].(float64)
	if !ok {
		http.Error(w, "Invalid humidity", http.StatusBadRequest)
		return
	}

	temperature.Set(temp)
	currentTemperature = temp

	humidity.Set(hum)
	currentHumidity = hum

	lastUpdated = time.Now()

	w.WriteHeader(http.StatusAccepted)
}
