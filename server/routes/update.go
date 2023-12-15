package routes

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/leomotors/home-env/services"
)

const (
	tempUpperBound = 85
	tempLowerBound = -40
	humUpperBound  = 100
	humLowerBound  = 0
)

func dataPostHandler(w http.ResponseWriter, r *http.Request) {
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

	sensorId, ok := data["sensorId"].(string)
	if !ok {
		http.Error(w, "Invalid sensorId", http.StatusBadRequest)
		return
	}

	sensorManager := services.GetSensorManager(sensorId)
	if sensorManager == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	temp, ok := data["temperature"].(float64)
	if !ok || math.IsNaN(temp) || temp < tempLowerBound || temp > tempUpperBound {
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	hum, ok := data["humidity"].(float64)
	if !ok || math.IsNaN(hum) || hum < humLowerBound || hum > humUpperBound {
		http.Error(w, "Invalid humidity", http.StatusBadRequest)
		return
	}

	sensorManager.SetValue(temp, hum)

	w.WriteHeader(http.StatusAccepted)
}

var UpdatePostHandler = http.HandlerFunc(dataPostHandler)
