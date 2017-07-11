package env

import (
	"fmt"
	"reflect"
	"strconv"
)

func handleUint(ref reflect.Value, structField reflect.StructField, rawVal string) error {
	if rawVal == "" {
		return nil
	}
	val, err := parseUint(structField, rawVal)
	if err != nil {
		return err
	}
	ref.SetUint(val)
	return nil
}

func parseUint(field reflect.StructField, rawVal string) (uint64, error) {
	size, err := getSize(field)

	i, err := strconv.ParseUint(rawVal, 10, size)
	if err != nil {
		return 0, err
	}

	// Get min/max values to check against
	min, err := getUintTag(field, "min", 0, size)
	if err != nil {
		return 0, err
	}
	max, err := getUintTag(field, "max", maxUints[size], size)
	if err != nil {
		return 0, err
	}

	if i < min {
		return 0, fmt.Errorf("%s must be at least %d", field.Name, min)
	}
	if i > max {
		return 0, fmt.Errorf("%s must be no more than %d", field.Name, max)
	}

	return i, nil
}

func getUintTag(structField reflect.StructField, tag string, defaultVal uint64, size int) (uint64, error) {
	rawVal, minExists := structField.Tag.Lookup(tag)
	if minExists {
		parsedVal, err := strconv.ParseUint(rawVal, 10, size)
		if err != nil {
			return 0, fmt.Errorf("unable to parse tag %s on %s: %s", tag, structField.Name, err)
		}
		return parsedVal, nil
	}
	return defaultVal, nil
}

func handleUintSlice(ref reflect.Value, structField reflect.StructField, rawArr []string) error {
	if len(rawArr) == 0 {
		return nil
	}
	var arr interface{}
	var err error

	t := ref.Type()
	switch t {
	case sliceOfUint8s:
		arr, err = getUint8Slice(structField, rawArr)
	case sliceOfUint16s:
		arr, err = getUint16Slice(structField, rawArr)
	case sliceOfUint32s:
		arr, err = getUint32Slice(structField, rawArr)
	case sliceOfUint64s:
		arr, err = getUint64Slice(structField, rawArr)
	case sliceOfUints:
		arr, err = getUintSlice(structField, rawArr)
	}

	if err != nil {
		return err
	}

	ref.Set(reflect.ValueOf(arr))
	return nil
}

func getUint8Slice(structField reflect.StructField, split []string) ([]uint8, error) {
	arr := make([]uint8, len(split), len(split))
	for i, raw := range split {
		val, err := parseUint(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = uint8(val)
	}
	return arr, nil
}

func getUint16Slice(structField reflect.StructField, split []string) ([]uint16, error) {
	arr := make([]uint16, len(split), len(split))
	for i, raw := range split {
		val, err := parseUint(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = uint16(val)
	}
	return arr, nil
}

func getUint32Slice(structField reflect.StructField, split []string) ([]uint32, error) {
	arr := make([]uint32, len(split), len(split))
	for i, raw := range split {
		val, err := parseUint(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = uint32(val)
	}
	return arr, nil
}

func getUint64Slice(structField reflect.StructField, split []string) ([]uint64, error) {
	arr := make([]uint64, len(split), len(split))
	for i, raw := range split {
		val, err := parseUint(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}

func getUintSlice(structField reflect.StructField, split []string) ([]uint, error) {
	arr := make([]uint, len(split), len(split))
	for i, raw := range split {
		val, err := parseUint(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = uint(val)
	}
	return arr, nil
}
