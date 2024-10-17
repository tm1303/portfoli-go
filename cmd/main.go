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

var isDebug bool = os.Getenv("ENV") == "DEBUG"

var targetCompanies []string = []string{}
var rootHoldingName string 
var inputFilename string

func initArgs(isDebug bool) {

	if isDebug {
		inputFilename = os.Args[1]
		rootHoldingName = os.Args[2]
		targetCompanies = os.Args[3:]
		return
	}
	inputFilename = os.Args[0]
	rootHoldingName = os.Args[1]
	targetCompanies = os.Args[2:]
}

func main() {

	initArgs(true)

	fmt.Println(inputFilename)
	fmt.Println(rootHoldingName)
	fmt.Println(targetCompanies)

	calcConfig := portfoligo.CalculatorConfig{
		RootHoldingName: rootHoldingName,
	}
	formatter := portfoligo.FormatterConfig{
		Target:           os.Stdout,
		HighlightHolding: targetCompanies,
	}

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
