package routes

import (
	"fmt"
	"net/http"

	"github.com/leomotors/home-env/constants"
	"github.com/leomotors/home-env/services"
)

// DataResponse represents the response for GET /data
type DataResponse struct {
	Temperature float64 `json:"temperature" example:"25.50" validate:"required"`
	Humidity    float64 `json:"humidity" example:"48.20" validate:"required"`
	LastUpdated float64 `json:"lastUpdated" example:"2.14" validate:"required"`
}

// @Summary		Get current sensor data
// @Description	Returns the latest temperature and humidity reading from the main room sensor
// @Produce		json
// @Success		200	{object}	DataResponse
// @Router			/data [get]
func publicDataGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sensorValue := services.GetSensorValue(constants.MainRoomId)

	data := fmt.Sprintf(
		`{"temperature": %.2f, "humidity": %.2f, "lastUpdated": %.2f}`,
		sensorValue.Temperature, sensorValue.Humidity, sensorValue.LastUpdated)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var DataGetHandler = http.HandlerFunc(publicDataGetHandler)
