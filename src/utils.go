package lighthouse

import (
	"net/http"
)

func getIPAddress(r *http.Request) string {
	return r.RemoteAddr
}