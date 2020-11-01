package fileio

import (
	excel "github.com/szyhf/go-excel"
)

func GetFromExcel(filePath string, config *excel.Config, model interface{}) error {
	conn := excel.NewConnecter()

	err := conn.Open(filePath)
	if err != nil {
		return err
	}
	defer conn.Close()

	rd, err := conn.NewReaderByConfig(config)
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
