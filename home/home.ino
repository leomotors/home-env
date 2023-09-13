#include <Adafruit_AHTX0.h>
#include <HTTPClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>

#include "password.h"

const char* ssid = WIFI_SSID;
const char* password = WIFI_PASSWORD;
const char* serverURL = SERVER_URL;
const char* serverPassword = SERVER_PASSWORD;

Adafruit_AHTX0 aht;

void setup() {
    Serial.begin(115200);
    Serial.println("Adafruit AHT10/AHT20 demo!");

    connectToWiFi();

    if (!aht.begin()) {
        Serial.println("Could not find AHT? Check wiring");
        while (true) delay(10);
    }

    Serial.println("AHT10 or AHT20 found");
}

void loop() {
    if (WiFi.status() != WL_CONNECTED) {
        // Wi-Fi is disconnected, attempt to reconnect
        connectToWiFi();
    }

    sensors_event_t humidity, temp;
    // populate temp and humidity objects with fresh data
    aht.getEvent(&humidity, &temp);
    Serial.print("Temperature: ");
    Serial.print(temp.temperature);
    Serial.println(" degrees C");
    Serial.print("Humidity: ");
    Serial.print(humidity.relative_humidity);
    Serial.println("% rH");

    updateData(humidity, temp);

    delay(3000);
}

void connectToWiFi() {
    Serial.println("Connecting to WiFi...");
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");
    }
    Serial.println("");
    Serial.print("Connected to WiFi network with IP Address: ");
    Serial.println(WiFi.localIP());
}

void updateData(sensors_event_t humidty, sensors_event_t temp) {
    WiFiClient client;
    HTTPClient http;

    http.begin(client, serverURL);
    http.addHeader("Content-Type", "application/json");
    http.addHeader("Authorization", serverPassword);

    String body = "{\"temperature\": " + String(temp.temperature) +
                  ", \"humidity\": " + String(humidty.relative_humidity) + "}";

    int responseCode = http.POST(body);

    Serial.print("HTTP Response code: ");
    Serial.println(responseCode);

    // Free resources
    http.end();
}
