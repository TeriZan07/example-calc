package calc

import (
	"context"
)

type Sum struct {
}

func (uc *Sum) Execute(ctx context.Context, numbers []float64) (result float64, count int) {
	for _, n := range numbers {
		result += n
		count += 1
	}

	return
}
