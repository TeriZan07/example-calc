package calc

import (
	"context"
)

type Mul struct {
}

func (uc *Mul) Execute(ctx context.Context, numbers []float64) (result float64, count int) {
	result = 1
	for _, n := range numbers {
		result *= n
		count += 1
	}

	return
}
