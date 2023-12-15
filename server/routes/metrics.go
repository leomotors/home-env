package routes

import "github.com/prometheus/client_golang/prometheus/promhttp"

var MetricsHandler = promhttp.Handler()
