package env

import (
	"fmt"
	"reflect"
	"strconv"
)

func handleFloat(ref reflect.Value, structField reflect.StructField, rawVal string, size int) error {
	if rawVal == "" {
		return nil
	}
	val, err := parseFloat(structField, rawVal, size)
	if err != nil {
		return err
	}
	ref.SetFloat(val)
	return nil
}

func parseFloat(structField reflect.StructField, rawVal string, size int) (float64, error) {
	f, err := strconv.ParseFloat(rawVal, size)
	if err != nil {
		return 0, err
	}

	// Get min/max values to check against
	min, err := getFloatTag(structField, "min", minFloats[size], size)
	if err != nil {
		return 0, err
	}
	max, err := getFloatTag(structField, "max", maxFloats[size], size)
	if err != nil {
		return 0, err
	}

	if f < min {
		return 0, fmt.Errorf("%s must be at least %f", structField.Name, min)
	}
	if f > max {
		return 0, fmt.Errorf("%s must be no more than %f", structField.Name, max)
	}

	return f, nil
}

func getFloatTag(structField reflect.StructField, tag string, defaultVal float64, size int) (float64, error) {
	rawVal, minExists := structField.Tag.Lookup(tag)
	if minExists {
		parsedVal, err := strconv.ParseFloat(rawVal, size)
		if err != nil {
			return 0, fmt.Errorf("unable to parse tag %s on %s: %s", tag, structField.Name, err)
		}
		return parsedVal, nil
	}
	return defaultVal, nil
}

func handleFloatSlice(ref reflect.Value, structField reflect.StructField, rawArr []string, size int) error {
	if len(rawArr) == 0 {
		return nil
	}
	var arr interface{}
	var err error

	switch size {
	case 32:
		arr, err = getFloat32Slice(structField, rawArr)
	case 64:
		arr, err = getFloat64Slice(structField, rawArr)
	}

	if err != nil {
		return err
	}

	ref.Set(reflect.ValueOf(arr))
	return nil
}

func getFloat32Slice(structField reflect.StructField, split []string) ([]float32, error) {
	arr := make([]float32, len(split), len(split))
	for i, raw := range split {
		val, err := parseFloat(structField, raw, 32)
		if err != nil {
			return arr, err
		}
		arr[i] = float32(val)
	}
	return arr, nil
}

func getFloat64Slice(structField reflect.StructField, split []string) ([]float64, error) {
	arr := make([]float64, len(split), len(split))
	for i, raw := range split {
		val, err := parseFloat(structField, raw, 32)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}
