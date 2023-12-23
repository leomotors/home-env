package middlewares

import (
	"net/http"
	"strconv"
	"strings"
)

func isLocalIP(ip string) bool {
	if ip[0:7] == "192.168" || ip[0:3] == "10." {
		return true
	}

	if ip[0:4] != "172." {
		return false
	}

	tokens := strings.Split(ip, ".")

	if len(tokens) != 4 {
		return false
	}

	i, _ := strconv.Atoi(tokens[1])

	return i >= 16 && i <= 31
}

func LocalOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		// Start with 172 or 192 only
		if !isLocalIP(ip) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
