package app

import (
	"github.com/unrolled/render"

	"skeleton/src/core/usecases/calc"
	"skeleton/src/handlers"
	"skeleton/src/handlers/api"
	"skeleton/src/handlers/front"
)

type Handlers struct {
	Index handlers.Handler
	Error handlers.Handler
	Sum   handlers.Handler
	Mul   handlers.Handler
}

func NewHandlers() Handlers {
	r := render.New(render.Options{
		Directory:     "src/front/views",
		Extensions:    []string{".html", ".tmpl"},
		Delims:        render.Delims{Left: "${", Right: "}"},
		IsDevelopment: true,
	})

	return Handlers{
		Index: &front.Index{Render: r},
		Error: &front.Error{Render: r},
		Sum:   &api.Calc{Render: r, CalcUC: &calc.Sum{}},
		Mul:   &api.Calc{Render: r, CalcUC: &calc.Mul{}},
	}
}
