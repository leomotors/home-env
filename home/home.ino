#include "config/const.hpp"
#include "config/password.hpp"
#include "services/aht.cpp"
#include "services/wifi.cpp"

void setup() {
    // Set pin mode for warning/check lights
    pinMode(config::WIFI_LED_PIN, OUTPUT);
    pinMode(config::AHT_LED_PIN, OUTPUT);
    pinMode(config::SERVER_LED_PIN, OUTPUT);

    // Turn on all warning/check lights, just like when starting car
    digitalWrite(config::WIFI_LED_PIN, HIGH);
    digitalWrite(config::AHT_LED_PIN, HIGH);
    digitalWrite(config::SERVER_LED_PIN, HIGH);

    Serial.begin(115200);
    Serial.println("Welcome to Home Environment Client!");

    delay(1000);
    connectToWiFi();
    waitForAHT();
}

void loop() {
    if (WiFi.status() != WL_CONNECTED) {
        // Wi-Fi is disconnected, attempt to reconnect
        connectToWiFi();
    }

    auto [success, value] = readSensor();

    if (success) {
        updateData(value.first, value.second);
    }

    delay(3000);
}
