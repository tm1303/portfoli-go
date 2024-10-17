package main

import (
	"bufio"
	"fmt"
	"os"
	"portfoli-go/portfoligo"
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

var targetCompanies []string = os.Args[3:]
var rootHoldingName string = os.Args[2]
var inputFilename string = os.Args[1]

func main() {

	//TODO: validate input args

	calcConfig := portfoligo.CalculatorConfig{
		RootHoldingName: rootHoldingName,
	}
	formatter := portfoligo.NewOrderedFormatter(
		portfoligo.FileWriterConfig{FileTarget: os.Stdout},
		targetCompanies,
	)

	p := portfoligo.NewPortfoligo(
		calcConfig,
		portfoligo.Load,
		formatter,
	)

	file, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, bufferSize)

	p.Run(reader)
}
