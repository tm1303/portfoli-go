package main

import (
	"bufio"
	"fmt"
	"os"
	"portfoli-go/portfligo"
)

// Assumtion:
// * the input file could be massive so lets be careful loading
// * "Ethical Global Fund" might not be the first fund in the file
// * "Ethical Global Fund" ~ "Global Ethical Fund" :)
// * there might be funds listed which are not in our portfolio
// * a fund is a holding with 1 or more holding
// * a copmany is a holding with exactly zero holdings
// * a holdings name won't be stupidly long, this is to tune our buffers
// * holding names are CASE SENSTAIVE (madness, but I probably won't have the energy for it)

var bufferSize int = 1024

// var rootHoldingName string = os.Args[2]
var inputFilename string = os.Args[1]

func main() {

	calcConfig := portfoligo.CalculatorConfig{
		RootHoldingName: os.Args[2],
	}

	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	reader := bufio.NewReaderSize(file, bufferSize)

	portfoligo.NewPortfoligo(calcConfig).Run(reader)

}
