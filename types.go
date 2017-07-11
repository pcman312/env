package env

import (
	"fmt"
	"net/url"
	"reflect"
	"time"
	"unsafe"
)

var (
	// Non-primitive types
	urlType = reflect.TypeOf(url.URL{})

	// Slice types
	sliceOfStrings = reflect.TypeOf([]string{})
	sliceOfBools   = reflect.TypeOf([]bool{})

	sliceOfInts   = reflect.TypeOf([]int{})
	sliceOfInt8s  = reflect.TypeOf([]int8{})
	sliceOfInt16s = reflect.TypeOf([]int16{})
	sliceOfInt32s = reflect.TypeOf([]int32{})
	sliceOfInt64s = reflect.TypeOf([]int64{})

	sliceOfUints   = reflect.TypeOf([]uint{})
	sliceOfUint8s  = reflect.TypeOf([]uint8{})
	sliceOfUint16s = reflect.TypeOf([]uint16{})
	sliceOfUint32s = reflect.TypeOf([]uint32{})
	sliceOfUint64s = reflect.TypeOf([]uint64{})

	sliceOfFloat32s = reflect.TypeOf([]float32{})
	sliceOfFloat64s = reflect.TypeOf([]float64{})

	sliceOfDurations = reflect.TypeOf([]time.Duration{})

	sliceOfUrlPointers = reflect.TypeOf([]*url.URL{})

	kindSizes = map[reflect.Kind]int{
		reflect.Int8:  int(unsafe.Sizeof(int8(0)) * 8),
		reflect.Int16: int(unsafe.Sizeof(int16(0)) * 8),
		reflect.Int32: int(unsafe.Sizeof(int32(0)) * 8),
		reflect.Int64: int(unsafe.Sizeof(int64(0)) * 8),
		reflect.Int:   int(unsafe.Sizeof(int(0)) * 8),

		reflect.Uint8:  int(unsafe.Sizeof(uint8(0)) * 8),
		reflect.Uint16: int(unsafe.Sizeof(uint16(0)) * 8),
		reflect.Uint32: int(unsafe.Sizeof(uint32(0)) * 8),
		reflect.Uint64: int(unsafe.Sizeof(uint64(0)) * 8),
		reflect.Uint:   int(unsafe.Sizeof(uint(0)) * 8),

		reflect.Float32: int(unsafe.Sizeof(float32(0)) * 8),
		reflect.Float64: int(unsafe.Sizeof(float64(0)) * 8),
	}
)

func getSize(field reflect.StructField) (int, error) {
	var sizeType reflect.Type // Type to actually analyze for size
	if field.Type.Kind() == reflect.Slice {
		sizeType = field.Type.Elem()
	} else {
		sizeType = field.Type
	}

	size, exists := kindSizes[sizeType.Kind()]
	if !exists {
		return 0, fmt.Errorf("invalid uint type: %s", field.Type.Kind())
	}
	return size, nil
}
