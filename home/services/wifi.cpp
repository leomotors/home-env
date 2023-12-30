#include <HTTPClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>

#include "../config/const.hpp"
#include "../config/password.hpp"

void connectToWiFi() {
    digitalWrite(config::WIFI_LED_PIN, HIGH);

    int retryCount = 0;
    const int maxRetryCount = 10;  // Maximum number of retry attempts

    Serial.println("Connecting to WiFi...");
    WiFi.begin(config::WIFI_SSID, config::WIFI_PASSWORD);

    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");

        retryCount++;
        if (retryCount >= maxRetryCount) {
            Serial.println(
                "Failed to connect to WiFi after multiple attempts. "
                "Retrying...");
            WiFi.begin(config::WIFI_SSID,
                       config::WIFI_PASSWORD);  // Retry connection
            retryCount = 0;                     // Reset retry count
        }
    }

    Serial.println("");
    Serial.print("Connected to WiFi network with IP Address: ");
    Serial.println(WiFi.localIP());

    digitalWrite(config::WIFI_LED_PIN, LOW);
}

void updateData(sensors_event_t humidty, sensors_event_t temp) {
    WiFiClient client;
    HTTPClient http;

    http.begin(client, config::SERVER_URL);
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Authorization", config::SERVER_PASSWORD);
    http.addHeader("X-Client-Version", String(config::CLIENT_VERSION));
    http.setUserAgent(config::USER_AGENT);

    String body = "{\"temperature\": " + String(temp.temperature) +
                  ", \"humidity\": " + String(humidty.relative_humidity) +
                  ", \"sensorId\": \"" + String(config::SENSOR_ID) + "\"}";

    int responseCode = http.POST(body);

    Serial.print("HTTP Response Code: ");
    Serial.println(responseCode);

    digitalWrite(config::SERVER_LED_PIN,
                 (responseCode >= 200 && responseCode < 300) ? LOW : HIGH);

    // Free resources
    http.end();
}
