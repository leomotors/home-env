package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	temperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature",
		Help: "Current temperature in degrees Celsius.",
	})
	humidity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_humidity",
		Help: "Current humidity level as a percentage.",
	})
	healthStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_health_status",
		Help: "ESP32 is maintaining connection with server.",
	})
)

var currentTemperature = 0.0
var currentHumidity = 0.0
var lastUpdated = time.Now()

func init() {
	prometheus.MustRegister(temperature)
	prometheus.MustRegister(humidity)
	prometheus.MustRegister(healthStatus)

	// Check if PASSWORD environment variable is set
	expectedPassword := os.Getenv("PASSWORD")
	if expectedPassword == "" {
		fmt.Println("ERROR: PASSWORD environment variable not set.")
		os.Exit(1)
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	if token == "" || channelID == "" {
		fmt.Println("ERROR: DISCORD_TOKEN or DISCORD_CHANNEL_ID environment variables not set.")
		os.Exit(1)
	}

	go setHealth()
}

func getLastUpdatedSeconds() float64 {
	return time.Since(lastUpdated).Seconds()
}

var alertThreshold = 10.0

func incrementLevel() {
	if alertThreshold >= 3600.0 {
		alertThreshold += 3600.0
		return
	}
	alertThreshold = alertLevelIncrement[alertThreshold]
}

var alertLevelIncrement = map[float64]float64{
	10.0:  60.0,
	60.0:  600.0,
	600.0: 3600.0,
}

func getAlertMessage() string {
	if alertThreshold > 3600.0 {
		return "2時間以上"
	}

	return alertLabel[alertThreshold]
}

var alertLabel = map[float64]string{
	10.0:   "10秒",
	60.0:   "1分",
	600.0:  "10分",
	3600.0: "1時間",
}

func setHealth() {
	for {
		lastUpdatedSeconds := getLastUpdatedSeconds()

		if lastUpdatedSeconds > 5 {
			healthStatus.Set(0)

			if lastUpdatedSeconds >= alertThreshold {
				fmt.Printf("Alerting: threshold = %f\n", alertThreshold)
				alertDiscord()
				incrementLevel()
			}
		} else {
			healthStatus.Set(1)
			alertThreshold = 10.0
		}
		time.Sleep(1 * time.Second)
	}
}

func alertDiscord() {
	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	apiUrl := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", channelID)

	requestData := map[string]interface{}{
		"content": fmt.Sprintf("# :warning::warning::warning: 市民请注意! :warning::warning::warning:\n## Your ESP32 does not do its job for %s!\n# CHECK IT NOW!\n如果您毫不犹豫、将从您的个人资料中扣除更多社会积分!!!", getAlertMessage()),
	}

	// Serialize the JSON data
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request with the JSON body
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header to indicate that the request body is in JSON format
	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with your token or credentials
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", token))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)

		var responseData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			fmt.Println("Error decoding JSON response:", err)
			return
		}

		fmt.Println("Response:", responseData)
		return
	}

	fmt.Println("Alert Success")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check authorization header
	expectedPassword := os.Getenv("PASSWORD")
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the HTML file
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	// Convert HTML content to string
	htmlString := string(htmlContent)

	// Replace a string in the HTML content
	replacedHTML := strings.Replace(
		htmlString,
		"{{ TEMPERATURE }}",
		fmt.Sprintf("%.2f", currentTemperature), -1)
	replacedHTML = strings.Replace(
		replacedHTML,
		"{{ HUMIDITY }}",
		fmt.Sprintf("%.2f", currentHumidity), -1)

	lastUpdatedSeconds := getLastUpdatedSeconds()
	replacedHTML = strings.Replace(
		replacedHTML,
		"{{ LAST_UPDATED }}",
		fmt.Sprintf("%.2f", lastUpdatedSeconds), -1)

	// Set the content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Serve the modified HTML content
	w.Write([]byte(replacedHTML))
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lastUpdatedSeconds := getLastUpdatedSeconds()
	data := fmt.Sprintf(
		"{\"temperature\": %.2f, \"humidity\": %.2f, \"lastUpdated\": %.2f}",
		currentTemperature, currentHumidity, lastUpdatedSeconds)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/update", postHandler)
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/data", getDataHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	fmt.Printf("Listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
