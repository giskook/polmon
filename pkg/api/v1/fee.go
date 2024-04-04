package v1

import "net/http"

func Fee(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			commonReply(w, http.StatusInternalServerError, "500000", nil, nil)
		}
	}()

	var http_status int
	var internal_status string
	var err error
	var data interface{}

	switch r.Method {
	case http.MethodGet:
		// http_status, internal_status, data, err = get(w, r)
	}

	commonReply(w, http_status, internal_status, data, err)
}
