package env

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

var (
	minDuration = time.Duration(math.MinInt64)
	maxDuration = time.Duration(math.MaxInt64)
)

func handleInt(value reflect.Value, field reflect.StructField, rawVal string, size int) error {
	if rawVal == "" {
		return nil
	}
	val, err := parseInt(field, rawVal, size)
	if err != nil {
		return err
	}
	value.SetInt(val)
	return nil
}

func parseInt(structField reflect.StructField, rawVal string, size int) (int64, error) {
	// Handle time.Duration parsing
	if structField.Type.String() == "time.Duration" {
		d, err := parseDuration(structField, rawVal)
		return d.Nanoseconds(), err
	}

	i, err := strconv.ParseInt(rawVal, 10, size)
	if err != nil {
		return 0, err
	}

	// Get min/max values to check against
	min, err := getIntTag(structField, "min", minInts[size], size)
	if err != nil {
		return 0, err
	}
	max, err := getIntTag(structField, "max", maxInts[size], size)
	if err != nil {
		return 0, err
	}

	if i < min {
		return 0, fmt.Errorf("%s must be at least %d", structField.Name, min)
	}
	if i > max {
		return 0, fmt.Errorf("%s must be no more than %d", structField.Name, max)
	}

	return i, nil
}

func getIntTag(structField reflect.StructField, tag string, defaultVal int64, size int) (int64, error) {
	rawVal, minExists := structField.Tag.Lookup(tag)
	if minExists {
		parsedVal, err := strconv.ParseInt(rawVal, 10, size)
		if err != nil {
			return 0, fmt.Errorf("unable to parse tag %s on %s: %s", tag, structField.Name, err)
		}
		return parsedVal, nil
	}
	return defaultVal, nil
}

func parseDuration(structField reflect.StructField, rawVal string) (time.Duration, error) {
	dur, err := time.ParseDuration(rawVal)
	if err != nil {
		return 0, err
	}

	// Get min/max values to check against
	min, err := getDurationTag(structField, "min", minDuration)
	if err != nil {
		return 0, err
	}
	max, err := getDurationTag(structField, "max", maxDuration)
	if err != nil {
		return 0, err
	}

	if dur < min {
		return 0, fmt.Errorf("%s must be at least %s", structField.Name, min)
	}
	if dur > max {
		return 0, fmt.Errorf("%s must be no more than %s", structField.Name, max)
	}

	return dur, nil
}

func getDurationTag(structField reflect.StructField, tag string, defaultVal time.Duration) (time.Duration, error) {
	rawVal, minExists := structField.Tag.Lookup(tag)
	if minExists {
		parsedVal, err := time.ParseDuration(rawVal)
		if err != nil {
			return 0, fmt.Errorf("unable to parse tag %s on %s: %s", tag, structField.Name, err)
		}
		return parsedVal, nil
	}
	return defaultVal, nil
}

func handleIntSlice(ref reflect.Value, structField reflect.StructField, rawArr []string, size int) error {
	if len(rawArr) == 0 {
		return nil
	}
	var arr interface{}
	var err error

	t := ref.Type()
	switch t {
	case sliceOfInt8s:
		arr, err = getInt8Slice(structField, rawArr)
	case sliceOfInt16s:
		arr, err = getInt16Slice(structField, rawArr)
	case sliceOfInt32s:
		arr, err = getInt32Slice(structField, rawArr)
	case sliceOfInt64s:
		arr, err = getInt64Slice(structField, rawArr)
	case sliceOfInts:
		arr, err = getIntSlice(structField, rawArr)
	case sliceOfDurations:
		arr, err = getDurationSlice(structField, rawArr)
	}

	if err != nil {
		return err
	}

	ref.Set(reflect.ValueOf(arr))
	return nil
}

func getInt8Slice(structField reflect.StructField, split []string) ([]int8, error) {
	arr := make([]int8, len(split), len(split))
	for i, raw := range split {
		val, err := parseInt(structField, raw, 8)
		if err != nil {
			return arr, err
		}
		arr[i] = int8(val)
	}
	return arr, nil
}

func getInt16Slice(structField reflect.StructField, split []string) ([]int16, error) {
	arr := make([]int16, len(split), len(split))
	for i, raw := range split {
		val, err := parseInt(structField, raw, 16)
		if err != nil {
			return arr, err
		}
		arr[i] = int16(val)
	}
	return arr, nil
}

func getInt32Slice(structField reflect.StructField, split []string) ([]int32, error) {
	arr := make([]int32, len(split), len(split))
	for i, raw := range split {
		val, err := parseInt(structField, raw, 32)
		if err != nil {
			return arr, err
		}
		arr[i] = int32(val)
	}
	return arr, nil
}

func getInt64Slice(structField reflect.StructField, split []string) ([]int64, error) {
	arr := make([]int64, len(split), len(split))
	for i, raw := range split {
		val, err := parseInt(structField, raw, 64)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}

func getIntSlice(structField reflect.StructField, split []string) ([]int, error) {
	arr := make([]int, len(split), len(split))
	for i, raw := range split {
		val, err := parseInt(structField, raw, 32)
		if err != nil {
			return arr, err
		}
		arr[i] = int(val)
	}
	return arr, nil
}

func getDurationSlice(structField reflect.StructField, rawArr []string) ([]time.Duration, error) {
	arr := make([]time.Duration, len(rawArr), len(rawArr))
	for i, raw := range rawArr {
		val, err := parseDuration(structField, raw)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}
