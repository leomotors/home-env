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

// UpdateRequest represents the request body for POST /update
type UpdateRequest struct {
	SensorID    string  `json:"sensorId" example:"main_room" validate:"required"`
	Temperature float64 `json:"temperature" example:"24.50" validate:"required"`
	Humidity    float64 `json:"humidity" example:"50.00" validate:"required"`
} //	@name	UpdateRequest

// @Summary		Update sensor data
// @Description	Receives temperature and humidity data from an ESP32 sensor. Local network only.
// @Accept			json
// @Param			Authorization	header	string			true	"Password for authentication"
// @Param			body			body	UpdateRequest	true	"Sensor reading data"
// @Success		202				"Accepted"
// @Failure		400				"Invalid request body or values out of range"
// @Failure		401				"Unauthorized"
// @Failure		404				"Sensor not found"
// @Router			/update [post]
func dataUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

var UpdatePostHandler = http.HandlerFunc(dataUpdateHandler)
