package helpers

import (
	"net/http"

	"github.com/go-chi/render"
)

type RequestError struct {
	Code        int    `json:"error"`
	Description string `json:"description"`
}

func (e *RequestError) Render(w http.ResponseWriter, r *http.Request) {
	render.Status(r, e.Code)
	render.JSON(w, r, e)
}
