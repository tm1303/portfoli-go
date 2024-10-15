package portfoligo

import (
	"fmt"
	"portfoli-go/domain"
)

var runningTotal float64 = float64(0)

type CalculatorConfig struct {
	RootHoldingName string
}

func (c CalculatorConfig) Calculate(holdingsIndex map[string]*domain.Holding) {
	rootHolding, ok := holdingsIndex[c.RootHoldingName]
	if !ok {
		panic("root holding not found")
	}

	step(rootHolding, float64(1))
}

func step(stepHolding *domain.Holding, stepWeight float64) {
	if len(stepHolding.Holdings) == 0 {
		runningTotal += stepWeight
		fmt.Printf("Company `%s` Weight `%f` (total: `%f`)\n", stepHolding.Name, stepWeight, runningTotal)
		return
	}

	for _, h := range stepHolding.Holdings {
		step(h.Holding, stepWeight*h.Weight)
	}
}
