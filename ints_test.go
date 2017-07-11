package env

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"

	"reflect"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_ints(t *testing.T) {
	Convey("int8", t, func() {
		defer resetEnv(os.Environ())
		testInt("myint8", 8)
	})

	Convey("int16", t, func() {
		defer resetEnv(os.Environ())
		testInt("myint16", 16)
	})

	Convey("int32", t, func() {
		defer resetEnv(os.Environ())
		testInt("myint32", 32)
	})

	Convey("int64", t, func() {
		defer resetEnv(os.Environ())
		testInt("myint64", 64)
	})

	Convey("int", t, func() {
		defer resetEnv(os.Environ())
		testInt("myint", 0)
	})

	Convey("int8 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyInt8Arr: []int8{1, 2, 3, 4},
		}

		os.Setenv("myint8arr", " 1, 2, 3, 4 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("int16 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyInt16Arr: []int16{5, 6, 7, 8},
		}

		os.Setenv("myint16arr", " 5, 6, 7, 8 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("int32 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyInt32Arr: []int32{-1, -2, -3, -4},
		}

		os.Setenv("myint32arr", " -1, -2, -3, -4 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("int64 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyInt64Arr: []int64{-5, -6, -7, -8},
		}

		os.Setenv("myint64arr", " -5, -6, -7, -8 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("int slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyIntArr: []int{1234, -290, 32, 4124},
		}

		os.Setenv("myintarr", " 1234, -290, 32, 4124 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("int8 min/max", t, func() {
		type TestStruct struct {
			MyInt8 int8 `env:"myint8" min:"8" max:"15"`
		}
		tests := map[int8]bool{
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
				MyInt8: value,
			}
			env := os.Environ()
			os.Setenv("myint8", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("int16 min/max", t, func() {
		type TestStruct struct {
			MyInt16 int16 `env:"myint16" min:"16" max:"31"`
		}
		tests := map[int16]bool{
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
				MyInt16: value,
			}
			env := os.Environ()
			os.Setenv("myint16", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("int32 min/max", t, func() {
		type TestStruct struct {
			MyInt32 int32 `env:"myint32" min:"32" max:"63"`
		}
		tests := map[int32]bool{
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
				MyInt32: value,
			}
			env := os.Environ()
			os.Setenv("myint32", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("int64 min/max", t, func() {
		type TestStruct struct {
			MyInt64 int64 `env:"myint64" min:"64" max:"127"`
		}
		tests := map[int64]bool{
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
				MyInt64: value,
			}
			env := os.Environ()
			os.Setenv("myint64", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("int min/max", t, func() {
		type TestStruct struct {
			MyInt int `env:"myint" min:"32" max:"63"`
		}
		tests := map[int]bool{
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
				MyInt: value,
			}
			env := os.Environ()
			os.Setenv("myint", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})
}

func TestGetIntTag(t *testing.T) {
	Convey("bad int tag", t, func() {
		defer resetEnv(os.Environ())

		type BadStructTag struct {
			MyField int `env:"myField" default:"asdf"`
		}

		bst := BadStructTag{}
		t := reflect.TypeOf(bst)
		field, ok := t.FieldByName("MyField")
		So(ok, ShouldBeTrue)

		_, err := getIntTag(field, "default", 1, 32)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unable to parse tag")
	})
}

func testInt(field string, size int) {
	rand.Seed(time.Now().UnixNano())
	var val int64
	expected := &TestConfig{}

	switch size {
	case 0:
		val = rand.Int63n(math.MaxInt32)
		expected.MyInt = int(val)
		os.Setenv(field, fmt.Sprint(val))
	case 8:
		val = rand.Int63n(math.MaxInt8)
		expected.MyInt8 = int8(val)
		os.Setenv(field, fmt.Sprint(val))
	case 16:
		val = rand.Int63n(math.MaxInt16)
		expected.MyInt16 = int16(val)
		os.Setenv(field, fmt.Sprint(val))
	case 32:
		val = rand.Int63n(math.MaxInt32)
		expected.MyInt32 = int32(val)
		os.Setenv(field, fmt.Sprint(val))
	case 64:
		val = rand.Int63n(math.MaxInt64)
		expected.MyInt64 = int64(val)
		os.Setenv(field, fmt.Sprint(val))
	default:
		panic("invalid size")
	}

	actual := &TestConfig{}

	err := Parse(actual)
	So(err, ShouldBeNil)
	So(actual, ShouldResemble, expected)
}
