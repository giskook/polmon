package v1

import "github.com/giskook/polmon/internal/persistence"

type Handler struct {
	persistence.Persistence
}

func NewHandlerV1(store persistence.Persistence) *Handler {
	return &Handler{
		Persistence: store,
	}
}
