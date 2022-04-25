package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func Csv2Json(csvPath string, jsonFilePath string, delimiter string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	}

	fields, err := reader.Read()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(jsonFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	defer f.Close()

	e := json.NewEncoder(f)
	e.SetIndent("", "  ")

	for {
		row, err := reader.Read()
		// continue to write data even if field count not match
		if errors.Is(err, io.EOF) {
			break
		} else if (err != nil) && !errors.Is(err, csv.ErrFieldCount) {
			return err
		}
		record := make(map[string]interface{})
		for idx, field := range fields {
			if idx < len(row) {
				record[field] = row[idx]
			}
		}
		e.Encode(record)
	}

	return nil
}
