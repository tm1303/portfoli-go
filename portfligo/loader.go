package portfoligo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

func Load(reader *bufio.Reader, output map[string]*domain.Holding) {

	decoder := json.NewDecoder(reader)
	_, err := decoder.Token()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading start of array: %v\n", err)
		return
	}

	for decoder.More() {
		var item holdingInput

		err := decoder.Decode(&item)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding JSON object: %v\n", err)
			return
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

	_, err = decoder.Token()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading end of array: %v\n", err)
		return
	}

}
