package test

import (
	"github.com/dinhtrung/xlsx-cli/pkg/format-converters/services"
	"testing"
)

func TestCSV2XLSX(t *testing.T) {
	if err := services.Csv2Xlsx("testdata/addresses.csv", "/tmp/addresses.xlsx", ","); err != nil {
		t.Fatal(err)
	}
}

func TestCSV2JSON(t *testing.T) {
	if err := services.Csv2Json("testdata/addresses.csv", "/tmp/addresses.json", ","); err != nil {
		t.Fatal(err)
	}
}
