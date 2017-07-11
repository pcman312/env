package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/pcman312/errutils"
)

var (
	ErrNotStructPointer = errors.New("input must be a pointer to a struct")
)

func Parse(conf interface{}) error {
	ptrRef := reflect.ValueOf(conf)
	if ptrRef.Kind() != reflect.Ptr {
		return ErrNotStructPointer
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotStructPointer
	}

	return parseStruct(ref)
}

func parseStruct(value reflect.Value) error {
	t := value.Type()
	errs := []error{}
	for i := 0; i < value.NumField(); i++ {
		err := handleField(value.Field(i), t.Field(i))
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errutils.JoinErrs(", ", errs...)
	}

	return nil
}

func handleField(value reflect.Value, field reflect.StructField) error {
	rawVal, err := getFieldValue(field)
	if err != nil {
		return err
	}

	err = parseField(value, field, rawVal)
	if err != nil {
		return err
	}

	return nil
}

func getFieldValue(field reflect.StructField) (string, error) {
	required, err := isRequired(field)
	if err != nil {
		return "", fmt.Errorf("'required' tag on %s must be either true or false\n", field.Name)
	}

	varName, hasEnv := field.Tag.Lookup("env")
	if !hasEnv || varName == "" || varName == "-" {
		fmt.Printf("Skipping field %s\n", field.Name)
		return "", nil
	}
	varName = strings.TrimSpace(varName)

	// Get value from environment
	rawValue := strings.TrimSpace(os.Getenv(varName))
	if rawValue != "" {
		return rawValue, nil
	}

	// No value in environment found, look at the default
	defaultVal := strings.TrimSpace(field.Tag.Get("default"))
	if defaultVal == "" && required {
		return "", fmt.Errorf("missing required variable [%s]", varName)
	}
	return defaultVal, nil
}

func isRequired(field reflect.StructField) (bool, error) {
	rawReq := strings.TrimSpace(field.Tag.Get("required"))
	if rawReq == "" {
		return false, nil
	}
	return strconv.ParseBool(rawReq)
}

func parseField(value reflect.Value, field reflect.StructField, rawVal string) error {
	switch field.Type.Kind() {
	case reflect.Bool:
		return handleBool(value, rawVal)

	case reflect.String:
		return handleString(value, rawVal)

	case reflect.Int8:
		return handleInt(value, field, rawVal, 8)
	case reflect.Int16:
		return handleInt(value, field, rawVal, 16)
	case reflect.Int, reflect.Int32:
		return handleInt(value, field, rawVal, 32)
	case reflect.Int64:
		return handleInt(value, field, rawVal, 64)

	case reflect.Uint8:
		return handleUint(value, field, rawVal, 8)
	case reflect.Uint16:
		return handleUint(value, field, rawVal, 16)
	case reflect.Uint, reflect.Uint32:
		return handleUint(value, field, rawVal, 32)
	case reflect.Uint64:
		return handleUint(value, field, rawVal, 64)

	case reflect.Float32:
		return handleFloat(value, field, rawVal, 32)
	case reflect.Float64:
		return handleFloat(value, field, rawVal, 64)

	case reflect.Slice:
		return handleSlice(value, field, rawVal)

	case reflect.Ptr:
		return handlePointer(value, field, rawVal)
	}

	return fmt.Errorf("unsupported type %s", field.Type.Kind())
}

func handleSlice(value reflect.Value, field reflect.StructField, rawVal string) error {
	delim := getSliceDelim(field)
	arr := make([]string, 0)
	if rawVal != "" {
		arr = strings.Split(rawVal, delim)

		for i, str := range arr {
			arr[i] = strings.TrimSpace(str)
		}
	}

	switch value.Type() {
	case sliceOfBools:
		return handleBoolSlice(value, arr)

	case sliceOfStrings:
		return handleStringSlice(value, arr)

	case sliceOfInt8s:
		return handleIntSlice(value, field, arr, 8)
	case sliceOfInt16s:
		return handleIntSlice(value, field, arr, 16)
	case sliceOfInt32s, sliceOfInts:
		return handleIntSlice(value, field, arr, 32)
	case sliceOfInt64s:
		return handleIntSlice(value, field, arr, 64)

	case sliceOfUint8s:
		return handleUintSlice(value, field, arr, 8)
	case sliceOfUint16s:
		return handleUintSlice(value, field, arr, 16)
	case sliceOfUint32s, sliceOfUints:
		return handleUintSlice(value, field, arr, 32)
	case sliceOfUint64s:
		return handleUintSlice(value, field, arr, 64)

	case sliceOfFloat32s:
		return handleFloatSlice(value, field, arr, 32)
	case sliceOfFloat64s:
		return handleFloatSlice(value, field, arr, 64)

	case sliceOfDurations:
		return handleIntSlice(value, field, arr, 64)

	case sliceOfUrlPointers:
		return handleUrlSlice(value, arr)

	default:
		return fmt.Errorf("unsupported slice type %s", field.Type.Elem().Kind())
	}
}

func getSliceDelim(field reflect.StructField) string {
	delim, delimExists := field.Tag.Lookup("delimiter")
	if !delimExists {
		delim = ","
	}

	return delim
}

func handlePointer(value reflect.Value, field reflect.StructField, rawVal string) error {
	switch field.Type.Elem() {
	case urlType:
		return handleUrl(value, rawVal)
	default:
		return fmt.Errorf("unsupported pointer type %s", field.Type.Elem().Kind())
	}
}
