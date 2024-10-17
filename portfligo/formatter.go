package portfoligo

import (
	"fmt"
	"os"
	"portfoli-go/domain"
	"slices"
	"sort"
)

type FormatterConfig struct {
	Target           *os.File
	HighlightHolding []string
}

func sortPortfolio(in map[string]domain.PortfolioLine) []domain.PortfolioLine {
	
	portfolioSlice := make([]domain.PortfolioLine, 0, len(in))

	for _, v := range in {
		portfolioSlice = append(portfolioSlice, v)
	}

	sort.Slice(portfolioSlice, func(i, j int) bool {
		return portfolioSlice[i].Weight > portfolioSlice[j].Weight
	})

	return portfolioSlice
}

const headingTag string = "\n     Company Name       :    Absolute Weight (%)\n    ---------------------------------------------\n"
const prefixTag string = "--->"
const emptyTag string = "\n   Portfolio contains no companies\n"


func (f FormatterConfig) Output(portflio map[string]domain.PortfolioLine) error {

	if len(portflio) == 0{
		if _, err := f.Target.WriteString(emptyTag); err != nil {
			return fmt.Errorf("failed to writing zero results: %w", err)
		}
		return nil
	}

	sorted := sortPortfolio(portflio)

	if _, err := f.Target.WriteString(headingTag); err != nil {
		return fmt.Errorf("failed to start writing portfolio: %w", err)
	}

	for _, v := range sorted {
		prefix := ""
		if slices.Index(f.HighlightHolding, v.Name) != -1 {
			prefix = prefixTag
		}
		if _, err := f.Target.WriteString(fmt.Sprintf("%-4s %-18s : %8.2f\n", prefix, v.Name, v.Weight*100)); err != nil {
			return fmt.Errorf("failed to write portfolio line: %w", err)
		}
	}

	return nil
}
