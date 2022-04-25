package services

import (
	"encoding/csv"
	"errors"
	"github.com/xuri/excelize/v2"
	"io"
	"os"
	"path"
)

func Csv2Xlsx(csvPath string, XLSXPath string, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	}
	xlsxFile := excelize.NewFile()
	sheetName := path.Base(csvPath)
	index := xlsxFile.NewSheet(sheetName)
	xlsxFile.SetActiveSheet(index)

	rowNum := 0
	for {
		row, err := reader.Read()
		// continue to write data even if field count not match
		if errors.Is(err, io.EOF) {
			break
		} else if (err != nil) && !errors.Is(err, csv.ErrFieldCount) {
			return err
		}
		rowNum++
		for colNum, cellValue := range row {
			cellName, err := excelize.CoordinatesToCellName(colNum+1, rowNum+1)
			if err != nil {
				return err
			}
			xlsxFile.SetCellValue(sheetName, cellName, cellValue)
		}
	}

	return xlsxFile.SaveAs(XLSXPath)
}
