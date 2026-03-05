-- Sensor readings (time-series data, replaces Prometheus gauges)
CREATE TABLE IF NOT EXISTS sensor_readings (
    time        TIMESTAMPTZ      NOT NULL,
    sensor_id   TEXT             NOT NULL,
    temperature DOUBLE PRECISION NOT NULL,
    humidity    DOUBLE PRECISION NOT NULL,
    is_migrated BOOLEAN          NOT NULL DEFAULT FALSE
);

SELECT create_hypertable('sensor_readings', 'time');

CREATE INDEX idx_sensor_readings_sensor_id_time
    ON sensor_readings (sensor_id, time DESC);

-- Downtime events (Grafana-friendly down/resolved pairs)
CREATE TABLE IF NOT EXISTS sensor_downtime_events (
    time        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    sensor_id   TEXT        NOT NULL,
    event_type  TEXT        NOT NULL CHECK (event_type IN ('down', 'resolved'))
);

SELECT create_hypertable('sensor_downtime_events', 'time');

CREATE INDEX idx_sensor_downtime_sensor_id_time
    ON sensor_downtime_events (sensor_id, time DESC);
