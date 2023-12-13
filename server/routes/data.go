package routes

import (
	"fmt"
	"net/http"

	"github.com/leomotors/home-env/constants"
	"github.com/leomotors/home-env/services"
)

func dataGetHandler(w http.ResponseWriter, r *http.Request) {
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

var DataGetHandler = http.HandlerFunc(dataGetHandler)
