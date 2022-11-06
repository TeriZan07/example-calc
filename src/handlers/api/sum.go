package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/unrolled/render"

	"skeleton/src/core/contracts"
	"skeleton/src/core/usecases/calc"
)

type Calc struct {
	Render *render.Render
	CalcUC calc.UseCase
}

func (h *Calc) Handle(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		_ = h.Render.Text(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	b, err := io.ReadAll(request.Body)
	if ok := HandleError(w, h.Render, err); !ok {
		return
	}

	var calcReq contracts.CalcReq
	err = json.Unmarshal(b, &calcReq)
	if ok := HandleError(w, h.Render, err); !ok {
		return
	}

	result, count := h.CalcUC.Execute(request.Context(), calcReq.Numbers)
	response := &contracts.CalcResp{Count: count, Result: result}
	err = h.Render.JSON(w, http.StatusOK, &response)
	HandleError(w, h.Render, err)
}
