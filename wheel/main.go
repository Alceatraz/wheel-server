package wheel

import (
	"log"
	"net/http"
	"wheel-server/wheel/healthz"
	"wheel-server/wheel/serverRoot"
)

func Entry() {

	mux := http.NewServeMux()

	healthz.AddHandlers(mux)
	serverRoot.AddHandlers(mux)

	server := &http.Server{Addr: ":80", Handler: mux}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalln("Error starting serverRoot:", err)
	}
}
