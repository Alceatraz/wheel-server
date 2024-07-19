package serverRoot

import "net/http"

func RegistHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handler)
}
