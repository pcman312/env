package env

import "math"

var (
	minInts = map[int]int64{
		8:  math.MinInt8,
		16: math.MinInt16,
		32: math.MinInt32,
		64: math.MinInt64,
	}
	maxInts = map[int]int64{
		8:  math.MaxInt8,
		16: math.MaxInt16,
		32: math.MaxInt32,
		64: math.MaxInt64,
	}

	maxUints = map[int]uint64{
		8:  math.MaxUint8,
		16: math.MaxUint16,
		32: math.MaxUint32,
		64: math.MaxUint64,
	}

	minFloats = map[int]float64{
		32: -1 * math.MaxFloat32,
		64: -1 * math.MaxFloat64,
	}
	maxFloats = map[int]float64{
		32: math.MaxFloat32,
		64: math.MaxFloat64,
	}
)
