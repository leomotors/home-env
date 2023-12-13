package routes

import "github.com/prometheus/client_golang/prometheus/promhttp"

var MetricsGetHandler = promhttp.Handler()
