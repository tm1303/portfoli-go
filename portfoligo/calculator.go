package portfoligo

import (
	"fmt"
	"portfoli-go/domain"
)

type CalculatorConfig struct {
	RootHoldingName string
}

func (c CalculatorConfig) Calculate(holdingsIndex map[string]*domain.Holding) (map[string]domain.PortfolioLine, error) {
	rootHolding, ok := holdingsIndex[c.RootHoldingName]
	if !ok {
		return nil, fmt.Errorf("root holding not found in holdings index")
	}

	portfolio := map[string]domain.PortfolioLine{}

	step(rootHolding, float64(1), portfolio)

	return portfolio, nil
}

func step(stepHolding *domain.Holding, stepWeight float64, portfolio map[string]domain.PortfolioLine) {
	if len(stepHolding.Holdings) == 0 {
		lineItem, ok := portfolio[stepHolding.Name]
		if !ok {
			lineItem = domain.PortfolioLine{
				Name: stepHolding.Name,
			}
		}
		lineItem.Weight += stepWeight
		portfolio[stepHolding.Name] = lineItem

		return
	}

	for _, h := range stepHolding.Holdings {
		step(h.Holding, stepWeight*h.Weight, portfolio)
	}
}
