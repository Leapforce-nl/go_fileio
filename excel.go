package fileio

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	excel "github.com/szyhf/go-excel"
)

func GetFromExcel(filePath string, config *excel.Config, model interface{}) *errortools.Error {
	conn := excel.NewConnecter()

	err := conn.Open(filePath)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	defer conn.Close()

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
