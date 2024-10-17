package portfoligo

import (
	"fmt"
	"os"
	"portfoli-go/domain"
	"slices"
	"sort"
)

//go:generate mockgen -source=formatter.go -destination=mocks/formatter_mocks.go -package=mocks
type TargetWriter interface {
	WriteInfoLine(s string) error
	WritePortfolioLine(name string, weight float64, highlight bool) error
}

// abstract away interacting with os.File
type FileWriterConfig struct {
	FileTarget *os.File
}

func (f FileWriterConfig) WriteInfoLine(s string) error {
	if _, err := f.FileTarget.WriteString(s); err != nil {
		return fmt.Errorf("failed to write info line: %w", err)
	}
	return nil
}

func (f FileWriterConfig) WritePortfolioLine(name string, weight float64, highlight bool) error {
	prefix := ""
	if highlight {
		prefix = prefixTag
	}
	if _, err := f.FileTarget.WriteString(fmt.Sprintf("%-4s %-18s : %8.2f\n", prefix, name, weight*100)); err != nil {
		return fmt.Errorf("failed to write info line: %w", err)
	}
	return nil
}

type orderedFileFormatterConfig struct {
	target           TargetWriter
	highlightHolding []string
}

type OrderedFileFormatter interface {
	Output(portflio map[string]domain.PortfolioLine) error
}

func NewOrderedFormatter(targetWriter TargetWriter, highlightHolding []string) OrderedFileFormatter {
	return orderedFileFormatterConfig{
		target:           targetWriter,
		highlightHolding: highlightHolding,
	}
}

func (f orderedFileFormatterConfig) Output(portflio map[string]domain.PortfolioLine) error {

	if len(portflio) == 0 {
		if err := f.target.WriteInfoLine(emptyTag); err != nil {
			return fmt.Errorf("failed to writing zero results: %w", err)
		}
		return nil
	}

	sorted := sortPortfolio(portflio)

	if err := f.target.WriteInfoLine(headingTag); err != nil {
		return fmt.Errorf("failed to start writing portfolio: %w", err)
	}

	for _, v := range sorted {
		highlight := slices.Index(f.highlightHolding, v.Name) != -1
		if err := f.target.WritePortfolioLine(v.Name, v.Weight, highlight); err != nil {
			return fmt.Errorf("failed to write portfolio line: %w", err)
		}
	}

	return nil
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
