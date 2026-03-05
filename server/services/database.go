package services

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB(databaseURL string) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	db = pool
	log.Println("Connected to database")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func InsertReading(sensorId string, temperature float64, humidity float64) {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO sensor_readings (time, sensor_id, temperature, humidity) VALUES ($1, $2, $3, $4)",
		time.Now(), sensorId, temperature, humidity,
	)
	if err != nil {
		log.Println("[DB] Failed to insert reading:", err)
	}
}

func InsertDowntimeEvent(sensorId string, eventType string) {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO sensor_downtime_events (time, sensor_id, event_type) VALUES ($1, $2, $3)",
		time.Now(), sensorId, eventType,
	)
	if err != nil {
		log.Println("[DB] Failed to insert downtime event:", err)
	}
}
