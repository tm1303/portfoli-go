package portfoligo

import (
	"bufio"
	"portfoli-go/domain"
)

type portfoligoConfig struct {
	calcConfig CalculatorConfig
}

type Portfoliogo interface {
	Run(reader *bufio.Reader)
}

func NewPortfoligo(calcConfig CalculatorConfig) Portfoliogo {
	return portfoligoConfig{
		calcConfig: calcConfig,
	}
}

func (p portfoligoConfig) Run(reader *bufio.Reader) {

	holdingsIndex := map[string]*domain.Holding{}
	Load(reader, holdingsIndex)
	p.calcConfig.Calculate(holdingsIndex)
}
