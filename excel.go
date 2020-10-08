package fileio

import (
	excel "github.com/szyhf/go-excel"
)

func GetFromExcel(filePath string, sheetName string, model interface{}) error {
	conn := excel.NewConnecter()

	err := conn.Open(filePath)
	if err != nil {
		return err
	}
	defer conn.Close()

	config := excel.Config{
		Sheet: sheetName,
	}
	rd, err := conn.NewReaderByConfig(&config)
	if err != nil {
		return err
	}
	defer rd.Close()

	err = rd.ReadAll(model)
	if err != nil {
		return err
	}

	return nil
}
