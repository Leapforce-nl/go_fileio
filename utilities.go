package fileio

import (
	"reflect"
	"strconv"

	types "github.com/Leapforce-nl/go_types"
)

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
				record = append(record, strconv.FormatInt(v1.Field(j).Int(), 10))
				break
			}
		}

		records = append(records, record)
	}

	return &records, nil
}
