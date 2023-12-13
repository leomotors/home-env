package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
		fmt.Println("[DISCORD ALERT] Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request with the JSON body
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("[DISCORD ALERT] Error creating request:", err)
		return
	}

	// Set the Content-Type header to indicate that the request body is in JSON format
	req.Header.Set("Content-Type", "application/json")

	// Set the Authorization header with your token or credentials
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", token))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[DISCORD ALERT] Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("[DISCORD ALERT] Request failed with status code:", resp.StatusCode)

		var responseData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			fmt.Println("[DISCORD ALERT] Error decoding JSON response:", err)
			return
		}

		fmt.Println("[DISCORD ALERT] Response:", responseData)
		return
	}

	fmt.Println("[DISCORD ALERT] Alert Success")
}
