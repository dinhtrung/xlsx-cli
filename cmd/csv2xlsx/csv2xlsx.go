package main

import (
	"flag"
	"github.com/dinhtrung/xlsx-cli/pkg/format-converters/services"
	"log"
)

var (
	inputFile, outputFile, delimiter string
)

// init declare run time flags for this application
func init() {
	flag.StringVar(&inputFile, "i", "", "input CSV file")
	flag.StringVar(&outputFile, "o", "", "output XLSX file")
	flag.StringVar(&delimiter, "d", ",", "delimiter character")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.Parse()
	if len(inputFile) == 0 {
		log.Fatal("invalid input file")
	}
	if len(outputFile) == 0 {
		log.Fatal("invalid output file")
	}
	err := services.Csv2Xlsx(inputFile, outputFile, delimiter)
	if err != nil {
		log.Fatal(err)
	}
}
