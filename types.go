package env

import (
	"net/url"
	"reflect"
	"time"
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
)
