package portfoligo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"portfoli-go/domain"
)

type subHoldingInput struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}
type holdingInput struct {
	Name        string            `json:"name"`
	SubHoldings []subHoldingInput `json:"holdings"`
}

func Load(reader *bufio.Reader) (map[string]*domain.Holding, error){

	decoder := json.NewDecoder(reader)
	_, err := decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("Error reading start of array: %w", err)
	}

	output := map[string]*domain.Holding{}
	for decoder.More() {
		var item holdingInput

		err := decoder.Decode(&item)
		if err != nil {
			return nil, fmt.Errorf ("Error decoding JSON object: %w", err)
		}

		h, ok := output[item.Name]
		if !ok {
			output[item.Name] = &domain.Holding{Name: item.Name}
			h = output[item.Name]
		}

		for _, s := range item.SubHoldings {
			subh, ok := output[s.Name]
			if !ok {
				output[s.Name] = &domain.Holding{Name: s.Name}
				subh = output[s.Name]
			}
			h.Holdings = append(h.Holdings, domain.SubHolding{
				Holding: subh,
				Weight:  s.Weight,
			})
		}
	}

	return output, nil
}
