package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func alertDiscord(msg string) {
	secret := GetSecret()
	token := secret.DISCORD_TOKEN
	channelID := secret.DISCORD_CHANNEL_ID

	apiUrl := fmt.Sprintf("https://discord.com/api/v10/channels/%d/messages", channelID)

	requestData := map[string]interface{}{
		"content": msg,
	}

	// Serialize the JSON data
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Println("[DISCORD ALERT] Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request with the JSON body
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("[DISCORD ALERT] Error creating request:", err)
		return
	}

	// Set the Content-Type header to indicate that the request body is in JSON format
	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with your token or credentials
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", token))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[DISCORD ALERT] Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Println("[DISCORD ALERT] Request failed with status code:", resp.StatusCode)

		var responseData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			log.Println("[DISCORD ALERT] Error decoding JSON response:", err)
			return
		}

		log.Println("[DISCORD ALERT] Response:", responseData)
		return
	}

	log.Println("[DISCORD ALERT] Alert Success")
}

var downAlertHeader = strings.Repeat("<:__mafuyuspook:1175812481373962331><:__kanadepolice:1184446548155838494>", 5)

func SendDownAlert(sensorId string, alertThreshold uint) {
	sensorManager := GetSensorManager(sensorId)

	timeText := GetAlertText(alertThreshold)
	sensorName := sensorManager.name

	if sensorName == "" {
		panic("Invalid Sensor Name! sensorId = " + sensorId)
	}

	alertText := fmt.Sprintf(
		`# %s
# :warning::warning::warning: 市民请注意! :warning::warning::warning:
## Your 小O (ESP32) 名=%s does not do its job for %s!
# CHECK IT NOW!
如果您毫不犹豫、将从您的个人资料中扣除更多社会积分!!!
# %s`, downAlertHeader, sensorName, timeText, overallStatusText())

	go alertDiscord(alertText)
}

var backNotiHeader = strings.Repeat("<:__honamiparty:1184457039125151764><:__mafuyuthunk:1184457025153945610>", 5)

func SendBackNotice(sensorId string) {
	sensorManager := GetSensorManager(sensorId)
	sensorName := sensorManager.name

	alertText := fmt.Sprintf(
		`# %s
# :white_check_mark::white_check_mark::white_check_mark: 干得好公民! :white_check_mark::white_check_mark::white_check_mark:
## Your 小O (ESP32) 名=%s is back to work!
# You are a good citizen!
没有共产党就没有新中国 没有共产党就没有新中国 !!!
# %s`, backNotiHeader, sensorName, overallStatusText())

	go alertDiscord(alertText)
}

func statusSymbol(status bool) string {
	if status {
		return ":white_check_mark:"
	} else {
		return ":warning:"
	}
}

func overallStatusText() string {
	sensorHealth := GetAllSensorHealth()

	statusTexts := make([]string, 0, len(*sensorHealth))

	for id, value := range *sensorHealth {
		statusTexts = append(statusTexts, fmt.Sprintf("[%s %s]", id, statusSymbol(value)))
	}

	return fmt.Sprintf("All Sensor Status: %s", strings.Join(statusTexts, " "))
}
