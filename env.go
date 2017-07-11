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

	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		return handleInt(value, field, rawVal)

	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return handleUint(value, field, rawVal)

	case reflect.Float32, reflect.Float64:
		return handleFloat(value, field, rawVal)

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

	case sliceOfInt8s, sliceOfInt16s, sliceOfInt32s, sliceOfInts, sliceOfInt64s, sliceOfDurations:
		return handleIntSlice(value, field, arr)

	case sliceOfUint8s, sliceOfUint16s, sliceOfUint32s, sliceOfUints, sliceOfUint64s:
		return handleUintSlice(value, field, arr)

	case sliceOfFloat32s, sliceOfFloat64s:
		return handleFloatSlice(value, field, arr)

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
