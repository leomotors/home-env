#include <Adafruit_AHTX0.h>

#include "../lib/pair.hpp"

Adafruit_AHTX0 aht;

void waitForAHT() {
    while (!aht.begin()) {
        Serial.println("Could not find AHT? Check wiring");
        delay(500);
    }

    Serial.println("AHT10 or AHT20 found");
}

CP::pair<bool, CP::pair<sensors_event_t, sensors_event_t>> readSensor() {
    sensors_event_t humidity, temp;

    // populate temp and humidity objects with fresh data
    const auto success = aht.getEvent(&humidity, &temp);

    // This does not work, getEvent will infinite loop if sensor is disconnected
    digitalWrite(config::AHT_LED_PIN, success ? LOW : HIGH);

    Serial.print("Temperature: ");
    Serial.print(temp.temperature);
    Serial.println(" °C");
    Serial.print("Humidity: ");
    Serial.print(humidity.relative_humidity);
    Serial.println("% rH");

    return {success, {humidity, temp}};
}
