package api

import "net/http"

type Handler interface {
	Fee(w http.ResponseWriter, r *http.Request)
}
