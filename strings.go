package env

import "reflect"

func handleString(value reflect.Value, rawVal string) error {
	value.SetString(rawVal)
	return nil
}

func handleStringSlice(value reflect.Value, rawArr []string) error {
	if len(rawArr) == 0 {
		return nil
	}
	value.Set(reflect.ValueOf(rawArr))
	return nil
}
