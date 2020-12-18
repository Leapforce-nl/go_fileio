package fileio

import (
	"encoding/csv"
	"io"
	"os"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

func GetCSVFromReader(reader io.Reader, model interface{}) *errortools.Error {
	csvReader := csv.NewReader(reader)

	return GetCSVFromCSVReader(csvReader, model)
}

func GetCSVFromCSVReader(reader *csv.Reader, model interface{}) *errortools.Error {
	records, err := reader.ReadAll()
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	e := utilities.StringArrayToStruct(&records, model)
	if e != nil {
		return e
	}

	return nil
}

func GetCSVFromFile(filePath string, model interface{}) *errortools.Error {
	file, err := os.Open(filePath)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	if file == nil {
		return nil
	}

	return GetCSVFromReader(file, model)
}

func WriteToCSV(filePath string, model interface{}, includeHeaders bool) *errortools.Error {
	file, err := os.Create(filePath)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return WriteToCSVFile(file, model, includeHeaders)
}

func WriteToCSVFile(file *os.File, model interface{}, includeHeaders bool) *errortools.Error {
	if file == nil {
		return nil
	}

	records, e := utilities.StructToStringArray(model, includeHeaders)
	if e != nil {
		return e
	}

	if records == nil {
		return nil
	}

	w := csv.NewWriter(file)
	err := w.WriteAll(*records) // calls Flush internally
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}
