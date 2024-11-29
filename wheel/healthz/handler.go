package healthz

import (
	"log"
	"net/http"
	"strconv"
)

var startupStatus = false
var livenessStatus = false
var readinessStatus = false

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleStartupStatus(w http.ResponseWriter, _ *http.Request) {
	if startupStatus {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleLivenessStatus(w http.ResponseWriter, _ *http.Request) {
	if livenessStatus {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleReadinessStatus(w http.ResponseWriter, _ *http.Request) {
	if readinessStatus {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func toggleStartupStatus(w http.ResponseWriter, _ *http.Request) {
	startupStatus = !startupStatus
	log.Println("Set /healthz/startup to -> " + strconv.FormatBool(startupStatus))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(strconv.FormatBool(startupStatus)))
}

func toggleLivenessStatus(w http.ResponseWriter, _ *http.Request) {
	livenessStatus = !livenessStatus
	log.Println("Set /healthz/liveness to -> " + strconv.FormatBool(startupStatus))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(strconv.FormatBool(startupStatus)))
}

func toggleReadinessStatus(w http.ResponseWriter, _ *http.Request) {
	readinessStatus = !readinessStatus
	log.Println("Set /healthz/readiness to -> " + strconv.FormatBool(readinessStatus))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(strconv.FormatBool(startupStatus)))
}
