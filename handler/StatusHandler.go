package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type StatusResponse struct {
	UptimeSeconds int `json:"uptime_seconds"`
}

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	//time.Since(startTime).Seconds()
	uptimeSeconds := int(time.Since(startTime).Seconds())

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(StatusResponse{
		UptimeSeconds: uptimeSeconds,
	})
}
