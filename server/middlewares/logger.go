package middlewares

import (
	"log"
	"net/http"

	"github.com/leomotors/home-env/utils"
)

func getIP(r *http.Request) string {
	ip := r.Header.Get("cf-connecting-ip")

	if ip == "" {
		ip = r.Header.Get("x-real-ip")
	}

	if ip == "" {
		ip = r.RemoteAddr
	}

	return ip
}

type statusCodeRecorder struct {
	http.ResponseWriter
	status int
}

func (recorder *statusCodeRecorder) WriteHeader(status int) {
	recorder.status = status
	recorder.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		method := r.Method
		path := r.URL.Path
		userAgent := r.Header.Get("User-Agent")

		recorder := &statusCodeRecorder{w, http.StatusOK}

		next.ServeHTTP(recorder, r)

		if path == "/metrics" && recorder.status == http.StatusOK {
			return
		}

		if path == "/update" && recorder.status == http.StatusAccepted {
			return
		}

		log.Printf("%s %s %s %d (%s)", ip, method, path, recorder.status, utils.TruncateString(userAgent, 30))
	})
}
