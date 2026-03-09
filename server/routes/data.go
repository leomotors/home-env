package routes

import (
	"encoding/json"
	"net/http"

	"github.com/leomotors/home-env/services"
)

// SensorDetail represents one sensor's data in the response
type SensorDetail struct {
	Id          string   `json:"id" example:"main_room" validate:"required"`
	Name        string   `json:"name" example:"Office Room" validate:"required"`
	Temperature *float64 `json:"temperature" example:"25.50"`
	Humidity    *float64 `json:"humidity" example:"48.20"`
	LastUpdated float64  `json:"lastUpdated" example:"2.14" validate:"required"`
	Online      bool     `json:"online" example:"true" validate:"required"`
} //	@name	SensorDetail

// DataResponse represents the response for GET /data
type DataResponse struct {
	Sensors []SensorDetail `json:"sensors" validate:"required"`
} //	@name	DataResponse

// @Summary		Get all sensor data
// @Description	Returns the latest temperature and humidity readings from all sensors
// @Produce		json
// @Success		200	{object}	DataResponse
// @Router			/data [get]
func publicDataGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	allSensors := services.GetAllSensors()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"sensors": allSensors}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var DataGetHandler = http.HandlerFunc(publicDataGetHandler)
