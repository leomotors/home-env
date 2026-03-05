package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/leomotors/home-env/constants"
	"github.com/leomotors/home-env/public"
	"github.com/leomotors/home-env/services"
)

func indexRenderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sensorValue := services.GetSensorValue(constants.MainRoomId)

	replacedHTML := strings.Replace(
		public.IndexHTML,
		"{{ TEMPERATURE }}",
		fmt.Sprintf("%.2f", sensorValue.Temperature), -1)
	replacedHTML = strings.Replace(
		replacedHTML,
		"{{ HUMIDITY }}",
		fmt.Sprintf("%.2f", sensorValue.Humidity), -1)
	replacedHTML = strings.Replace(
		replacedHTML,
		"{{ LAST_UPDATED }}",
		fmt.Sprintf("%.2f", sensorValue.LastUpdated), -1)

	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Serve the modified HTML content
	w.Write([]byte(replacedHTML))
}

var IndexGetHandler = http.HandlerFunc(indexRenderHandler)
