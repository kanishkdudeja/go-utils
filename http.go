package utils

import (
	"net"
	"net/http"
	"strings"
)

/* The following functions are for finding the correct IP of the visitor. It takes into account
if a reverse proxy has been setup (forwards X-REAL-IP in that case). It also takes into
account that a user might be using a proxy itself */

// GetIPAdress function tries to get the remote ip and returns it or ""
func GetIPAdress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			}
			return ip
		}
	}
	return ""
}

// GetRemoteIP function tries to get the remote ip and returns it or ""
func GetRemoteIP(r *http.Request) string {
	a := r.Header.Get("X-Real-IP")

	if a == "" {
		a = r.Header.Get("X-Forwarded-For")
	}

	if a == "" {
		a = strings.SplitN(r.RemoteAddr, ":", 2)[0]

		// Check localhost
		if a == "[" {
			a = "127.0.0.1"
		}
	}

	return a
}

// GetUserAgent function returns the UserAgent for the supplied http.Request
func GetUserAgent(r *http.Request) string {
	return r.UserAgent()
}
