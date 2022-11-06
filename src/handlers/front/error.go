package front

import (
	"net/http"

	"github.com/unrolled/render"
)

type Error struct {
	Render *render.Render
}

func (h *Error) Handle(w http.ResponseWriter, r *http.Request) {
	_ = h.Render.HTML(w, http.StatusInternalServerError, "error/error", ErrorBind())
}
