package healthz

import "net/http"

func AddHandlers(mux *http.ServeMux) {

	mux.HandleFunc("GET /healthz", healthz)

	mux.HandleFunc("GET /healthz/startup", handleStartupStatus)
	mux.HandleFunc("GET /healthz/liveness", handleLivenessStatus)
	mux.HandleFunc("GET /healthz/readiness", handleReadinessStatus)

	mux.HandleFunc("POST /healthz/startup", toggleStartupStatus)
	mux.HandleFunc("POST /healthz/liveness", toggleLivenessStatus)
	mux.HandleFunc("POST /healthz/readiness", toggleReadinessStatus)

}
