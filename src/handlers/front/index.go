package front

import (
	"net/http"

	"github.com/unrolled/render"
)

type Index struct {
	Render *render.Render
}

func (h *Index) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_ = h.Render.HTML(w, http.StatusOK, "index/index", CommonBind())
}
