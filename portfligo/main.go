package portfoligo

import (
	"bufio"
	"fmt"
	"portfoli-go/domain"
)

type loaderFunc func(reader *bufio.Reader) (map[string]*domain.Holding, error)

type PortfolioWriter interface
{ 
	Output(portflio map[string]domain.PortfolioLine) error
}

type Calculator interface{
	Calculate(holdingsIndex map[string]*domain.Holding) (map[string]domain.PortfolioLine, error)
}

type PortfoligoConfig struct {
	calc Calculator
	loaderFunc
	writer PortfolioWriter
}

func NewPortfoligo(calcConfig CalculatorConfig, loaderFunc loaderFunc, portfolioWriter PortfolioWriter) PortfoligoConfig {
	return PortfoligoConfig{
		calc: calcConfig,
		loaderFunc: loaderFunc,
		writer: portfolioWriter,
	}
}

func (p PortfoligoConfig) Run(reader *bufio.Reader) {
	holdingsIndex, err := p.loaderFunc(reader)
	if err != nil {
		panic (fmt.Errorf ("error loading holdings: %w", err))
	}
	companiesWeighted, err := p.calc.Calculate(holdingsIndex)
	if err != nil {
		panic (fmt.Errorf ("error calculating company weights: %w", err))
	}
	err = p.writer.Output(companiesWeighted)
	if err != nil {
		panic (fmt.Errorf ("error writing to output: %w", err))
	}
}
