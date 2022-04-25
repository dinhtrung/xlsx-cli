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
	flag.StringVar(&inputFile, "test", "", "input CSV file")
	flag.StringVar(&outputFile, "skip-api", "", "output XLSX file")
	flag.StringVar(&delimiter, ",", ",", "delimiter character")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.Parse()
	err := services.Csv2Json(inputFile, outputFile, delimiter)
	if err != nil {
		log.Fatal(err)
	}
}
