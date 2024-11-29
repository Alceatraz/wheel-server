package serverRoot

import "net/http"

func AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handler)
}
