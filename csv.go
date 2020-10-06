package fileio

import (
	"encoding/csv"
	"os"
)

func GetFromCSV(filePath string, sheetName string, model interface{}) (*[][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	if file == nil {
		return nil, nil
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &records, nil
}

func WriteToCSV(filePath string, model interface{}, includeHeaders bool) error {
	records, err := StructToStringArray(model, includeHeaders)
	if err != nil {
		return err
	}

	if records == nil {
		return nil
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(*records) // calls Flush internally
	if err != nil {
		return err
	}

	return nil
}
