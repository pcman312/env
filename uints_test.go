package env

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_uints(t *testing.T) {
	Convey("uint8", t, func() {
		defer resetEnv(os.Environ())

		testUint("myuint8", 8)
	})

	Convey("uint16", t, func() {
		defer resetEnv(os.Environ())

		testUint("myuint16", 16)
	})

	Convey("uint32", t, func() {
		defer resetEnv(os.Environ())

		testUint("myuint32", 32)
	})

	Convey("uint64", t, func() {
		defer resetEnv(os.Environ())

		testUint("myuint64", 64)
	})

	Convey("uint", t, func() {
		defer resetEnv(os.Environ())

		testUint("myuint", 0)
	})

	Convey("uint8 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyUint8Arr: []uint8{1, 2, 3, 4},
		}

		os.Setenv("myuint8arr", " 1, 2, 3, 4 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("uint16 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyUint16Arr: []uint16{5, 6, 7, 8},
		}

		os.Setenv("myuint16arr", " 5, 6, 7, 8 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("uint32 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyUint32Arr: []uint32{9, 10, 11, 12},
		}

		os.Setenv("myuint32arr", " 9, 10, 11, 12 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("uint64 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyUint64Arr: []uint64{13, 14, 15, 16},
		}

		os.Setenv("myuint64arr", " 13, 14, 15, 16 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("uint slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyUintArr: []uint{1234, 290, 32, 4124},
		}

		os.Setenv("myuintarr", " 1234, 290, 32, 4124 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("uint8 min/max", t, func() {
		type TestStruct struct {
			MyUint8 uint8 `env:"myuint8" min:"8" max:"15"`
		}
		tests := map[uint8]bool{
			6:  false,
			7:  false,
			8:  true,
			9:  true,
			14: true,
			15: true,
			16: false,
			17: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyUint8: value,
			}
			env := os.Environ()
			os.Setenv("myuint8", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("uint16 min/max", t, func() {
		type TestStruct struct {
			MyUint16 uint16 `env:"myuint16" min:"16" max:"31"`
		}
		tests := map[uint16]bool{
			14: false,
			15: false,
			16: true,
			17: true,
			30: true,
			31: true,
			32: false,
			33: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyUint16: value,
			}
			env := os.Environ()
			os.Setenv("myuint16", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("uint32 min/max", t, func() {
		type TestStruct struct {
			MyUint32 uint32 `env:"myuint32" min:"32" max:"63"`
		}
		tests := map[uint32]bool{
			30: false,
			31: false,
			32: true,
			33: true,
			62: true,
			63: true,
			64: false,
			65: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyUint32: value,
			}
			env := os.Environ()
			os.Setenv("myuint32", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("uint64 min/max", t, func() {
		type TestStruct struct {
			MyUint64 uint64 `env:"myuint64" min:"64" max:"127"`
		}
		tests := map[uint64]bool{
			62:  false,
			63:  false,
			64:  true,
			65:  true,
			126: true,
			127: true,
			128: false,
			129: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyUint64: value,
			}
			env := os.Environ()
			os.Setenv("myuint64", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("uint min/max", t, func() {
		type TestStruct struct {
			MyUint uint `env:"myuint" min:"32" max:"63"`
		}
		tests := map[uint]bool{
			30: false,
			31: false,
			32: true,
			33: true,
			62: true,
			63: true,
			64: false,
			65: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyUint: value,
			}
			env := os.Environ()
			os.Setenv("myuint", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})
}

func TestGetUintTag(t *testing.T) {
	Convey("bad uint tag", t, func() {
		defer resetEnv(os.Environ())

		type BadStructTag struct {
			MyField uint `env:"myField" default:"asdf"`
		}

		bst := BadStructTag{}
		t := reflect.TypeOf(bst)
		field, ok := t.FieldByName("MyField")
		So(ok, ShouldBeTrue)

		_, err := getUintTag(field, "default", 1, 32)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unable to parse tag")
	})
}

func testUint(field string, size int) {
	rand.Seed(time.Now().UnixNano())
	expected := &TestConfig{}

	switch size {
	case 0:
		val := getRandUint(32)
		expected.MyUint = uint(val)
		os.Setenv(field, fmt.Sprint(val))
	case 8:
		val := getRandUint(8)
		expected.MyUint8 = uint8(val)
		os.Setenv(field, fmt.Sprint(val))
	case 16:
		val := getRandUint(16)
		expected.MyUint16 = uint16(val)
		os.Setenv(field, fmt.Sprint(val))
	case 32:
		val := getRandUint(32)
		expected.MyUint32 = uint32(val)
		os.Setenv(field, fmt.Sprint(val))
	case 64:
		val := getRandUint(64)
		expected.MyUint64 = uint64(val)
		os.Setenv(field, fmt.Sprint(val))
	default:
		panic("invalid size")
	}

	actual := &TestConfig{}

	err := Parse(actual)
	So(err, ShouldBeNil)
	So(actual, ShouldResemble, expected)
}

func getRandUint(size int) uint64 {
	var val uint64
	val = rand.Uint64()
	switch size {
	case 8:
		val %= math.MaxUint8
	case 16:
		val %= math.MaxUint16
	case 0, 32:
		val %= math.MaxUint32
	case 64:
		val %= math.MaxUint64
	default:
		panic("invalid size")
	}
	return val
}
