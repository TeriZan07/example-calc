package calc

import "context"

type UseCase interface {
	Execute(context.Context, []float64) (float64, int)
}
