package http

import (
	"net/http"
	"strings"
)

// Page is
type Page struct {
	Title string
	Body  []byte
}

// Handler is a collection of all the service handlers.
type Handler struct {
	Page *Page
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/dials") {
		//	h.Handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}
