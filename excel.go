package fileio

import (
	"io"
	"io/ioutil"

	errortools "github.com/leapforce-libraries/go_errortools"
	excel "github.com/szyhf/go-excel"
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
