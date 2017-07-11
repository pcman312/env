package env

import (
	"fmt"
	"reflect"
	"strconv"
)

func handleFloat(ref reflect.Value, structField reflect.StructField, rawVal string) error {
	if rawVal == "" {
		return nil
	}
	val, err := parseFloat(structField, rawVal)
	if err != nil {
		return err
	}
	ref.SetFloat(val)
	return nil
}

func parseFloat(field reflect.StructField, rawVal string) (float64, error) {
	size, err := getSize(field)
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(rawVal, size)
	if err != nil {
		return 0, err
	}

	// Get min/max values to check against
	min, err := getFloatTag(field, "min", minFloats[size], size)
	if err != nil {
		return 0, err
	}
	max, err := getFloatTag(field, "max", maxFloats[size], size)
	if err != nil {
		return 0, err
	}

	if f < min {
		return 0, fmt.Errorf("%s must be at least %f", field.Name, min)
	}
	if f > max {
		return 0, fmt.Errorf("%s must be no more than %f", field.Name, max)
	}

	return f, nil
}

func getFloatTag(field reflect.StructField, tag string, defaultVal float64, size int) (float64, error) {
	rawVal, minExists := field.Tag.Lookup(tag)
	if minExists {
		parsedVal, err := strconv.ParseFloat(rawVal, size)
		if err != nil {
			return 0, fmt.Errorf("unable to parse tag %s on %s: %s", tag, field.Name, err)
		}
		return parsedVal, nil
	}
	return defaultVal, nil
}

func handleFloatSlice(ref reflect.Value, field reflect.StructField, rawArr []string) error {
	if len(rawArr) == 0 {
		return nil
	}
	var arr interface{}
	var err error

	t := ref.Type()
	switch t {
	case sliceOfFloat32s:
		arr, err = getFloat32Slice(field, rawArr)
	case sliceOfFloat64s:
		arr, err = getFloat64Slice(field, rawArr)
	}

	if err != nil {
		return err
	}

	ref.Set(reflect.ValueOf(arr))
	return nil
}

func getFloat32Slice(field reflect.StructField, split []string) ([]float32, error) {
	arr := make([]float32, len(split), len(split))
	for i, raw := range split {
		val, err := parseFloat(field, raw)
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
		val, err := parseFloat(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}
