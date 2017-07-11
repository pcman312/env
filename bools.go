package env

import (
	"reflect"
	"strconv"
)

func handleBool(value reflect.Value, rawVal string) error {
	val, err := parseBool(rawVal)
	if err != nil {
		return err
	}
	value.SetBool(val)
	return nil
}

func parseBool(rawVal string) (bool, error) {
	if rawVal == "" {
		return false, nil
	}
	return strconv.ParseBool(rawVal)
}

func handleBoolSlice(value reflect.Value, rawArr []string) error {
	if len(rawArr) == 0 {
		return nil
	}
	arr := make([]bool, len(rawArr), len(rawArr))
	for i, rawVal := range rawArr {
		val, err := parseBool(rawVal)
		if err != nil {
			return err
		}
		arr[i] = val
	}

	value.Set(reflect.ValueOf(arr))
	return nil
}
