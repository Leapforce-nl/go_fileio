package fileio

import (
	"reflect"
	"strconv"

	types "github.com/Leapforce-nl/go_types"
)

func StringArrayToStruct(records *[][]string, model interface{}) error {
	if records == nil {
		return nil
	}

	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return &types.ErrorString{"The interface is not a pointer."}
	}

	v := reflect.ValueOf(model).Elem()
	if v.Kind() != reflect.Slice {
		return &types.ErrorString{"The interface is not a pointer to a slice."}
	}

	rv := reflect.ValueOf(model)

	structType := reflect.TypeOf(model).Elem().Elem()

	numFields := structType.NumField()

	fields := make(map[string]int)

	for index, record := range *records {
		if index == 0 {
			for cellIndex, cellValue := range record {
				fields[cellValue] = cellIndex
			}

			continue
		}

		new := reflect.New(structType).Elem()

		for i := 0; i < numFields; i++ {
			fieldName := structType.Field(i).Name
			fieldTag := structType.Field(i).Tag.Get("csv")

			if fieldTag == "" {
				continue
			}
			fieldIndex, ok := fields[fieldTag]

			if ok {
				value := record[fieldIndex]

				switch new.FieldByName(fieldName).Kind() {
				case reflect.String:
					new.FieldByName(fieldName).SetString(value)
					break
				case reflect.Int:
					i, err := strconv.ParseInt(value, 10, 64)
					if err == nil {
						new.FieldByName(fieldName).SetInt(i)
					}
					break
				}

			}
		}

		rv.Elem().Set(reflect.Append(rv.Elem(), new))
	}

	return nil
}

func StructToStringArray(model interface{}, includeHeaders bool) (*[][]string, error) {
	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Slice {
		return nil, &types.ErrorString{"The interface is not a slice."}
	}

	records := [][]string{}

	if includeHeaders {
		e := reflect.TypeOf(model).Elem()
		record := []string{}
		for i := 0; i < e.NumField(); i++ {
			record = append(record, e.Field(i).Name)
		}

		records = append(records, record)
	}

	for i := 0; i < v.Len(); i++ {

		record := []string{}
		v1 := v.Index(i)
		for j := 0; j < v1.NumField(); j++ {
			switch v1.Field(j).Kind() {
			case reflect.String:
				record = append(record, v1.Field(j).String())
				break
			case reflect.Int:
				record = append(record, strconv.FormatInt(v1.Field(j).Int(), 10))
				break
			default:
				record = append(record, "")
				break
			}
		}

		records = append(records, record)
	}

	return &records, nil
}
