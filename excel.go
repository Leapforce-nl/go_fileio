package fileio

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"

	errortools "github.com/leapforce-libraries/go_errortools"
	excel "github.com/leapforce-libraries/go_excel"
)

func GetFromExcelFile(filePath string, config *excel.Config, model interface{}) *errortools.Error {
	conn := excel.NewConnecter()

	err := conn.Open(filePath)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	defer conn.Close()

	return readExcel(conn, config, model)
}

func GetFromExcelReader(reader io.Reader, config *excel.Config, model interface{}) *errortools.Error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	conn := excel.NewConnecter()

	err = conn.OpenBinary(b)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	defer conn.Close()

	return readExcel(conn, config, model)
}

func readExcel(conn excel.Connecter, config *excel.Config, model interface{}) *errortools.Error {
	rd, err := conn.NewReaderByConfig(config)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	defer rd.Close()

	err = rd.ReadAll(model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}

func ParseExcelDate(value string) (*civil.Date, *errortools.Error) {
	value = strings.Trim(value, " ")

	if value == "" {
		return nil, nil
	}

	t, err := time.Parse("02-01-2006", value)
	if err != nil {
		t, err = time.Parse("2-1-2006", value)
	}
	if err == nil {
		date := civil.DateOf(t)
		return &date, nil
	}

	dateFloat64, err := strconv.ParseFloat(value, 64)

	if err == nil {
		// Excel "day 0" is 1899-12-30
		t0, _ := time.Parse("2006-01-02", "1899-12-30")
		date := civil.DateOf(t0.Add(time.Duration(int64(dateFloat64)*24) * time.Hour))
		return &date, nil
	}

	return nil, errortools.ErrorMessage(fmt.Sprintf("Cannot convert '%s' to civil.Date", value))
}
