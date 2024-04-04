package v1

import "net/http"

func (h *Handler) Fee(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			commonReply(w, http.StatusInternalServerError, "500000", nil, nil)
		}
	}()

	var httpStatus int
	var internalStatus string
	var err error
	var data interface{}

	switch r.Method {
	case http.MethodGet:
		fee, err := h.GetTotalFee()
		if err != nil {
			httpStatus = http.StatusInternalServerError
			internalStatus = "500001"
			data = nil
		} else {
			httpStatus = http.StatusOK
			internalStatus = "200000"
			data = fee
		}
	}

	commonReply(w, httpStatus, internalStatus, data, err)
}
