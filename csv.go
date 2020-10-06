package fileio

import (
	"encoding/csv"
	"os"
)

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
