package api

import (
	"net/http"

	"github.com/unrolled/render"
)

func HandleError(w http.ResponseWriter, r *render.Render, err error) (ok bool) {
	if err != nil {
		_ = r.Text(w, http.StatusInternalServerError, err.Error())
		return false
	}

	return true
}
