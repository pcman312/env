package env

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_floats(t *testing.T) {
	Convey("float32", t, func() {
		defer resetEnv(os.Environ())

		testFloat("myfloat32", 32)
	})

	Convey("float64", t, func() {
		defer resetEnv(os.Environ())

		testFloat("myfloat64", 64)
	})

	Convey("float32 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyFloat32Arr: []float32{2345.51, 7452.511, 713.12},
		}

		os.Setenv("myfloat32arr", " 2345.51, 7452.511, 713.12 ")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("float64 slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}

		// Be careful updating these floats because of rounding errors
		floats := []float64{
			2345.199951171875,
			7452.2001953125,
			713.2000122070312,
			math.MaxFloat32 + 1,
		}
		strs := make([]string, len(floats), len(floats))
		for i, f := range floats {
			strs[i] = fmt.Sprint(f)
		}

		expected := &TestConfig{
			MyFloat64Arr: floats,
		}

		os.Setenv("myfloat64arr", strings.Join(strs, ", "))
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("float32 min/max", t, func() {
		type TestStruct struct {
			MyFloat32 float32 `env:"myfloat32" min:"500.25" max:"999.75"`
		}
		tests := map[float32]bool{
			499.0:  false,
			500.0:  false,
			501.0:  true,
			502.0:  true,
			998.0:  true,
			999.0:  true,
			1000.0: false,
			1001.0: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyFloat32: value,
			}
			env := os.Environ()
			os.Setenv("myfloat32", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})

	Convey("float64 min/max", t, func() {
		type TestStruct struct {
			MyFloat64 float64 `env:"myfloat64" min:"1500.25" max:"1999.75"`
		}
		tests := map[float64]bool{
			1499.0: false,
			1500.0: false,
			1501.0: true,
			1502.0: true,
			1998.0: true,
			1999.0: true,
			2000.0: false,
			2001.0: false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}
			expected := &TestStruct{
				MyFloat64: value,
			}
			env := os.Environ()
			os.Setenv("myfloat64", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})
}

func TestGetFloatTag(t *testing.T) {
	Convey("bad float tag", t, func() {
		defer resetEnv(os.Environ())

		type BadStructTag struct {
			MyField float64 `env:"myField" default:"asdf"`
		}

		bst := BadStructTag{}
		t := reflect.TypeOf(bst)
		field, ok := t.FieldByName("MyField")
		So(ok, ShouldBeTrue)

		_, err := getFloatTag(field, "default", 1, 64)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unable to parse tag")
	})
}

func testFloat(field string, size int) {
	rand.Seed(time.Now().UnixNano())
	var val float64
	expected := &TestConfig{}

	switch size {
	case 32:
		val = getRandFloat(32)
		expected.MyFloat32 = float32(val)
		os.Setenv(field, fmt.Sprint(val))
	case 64:
		val = getRandFloat(64)
		expected.MyFloat64 = float64(val)
		os.Setenv(field, fmt.Sprint(val))
	}

	actual := &TestConfig{}

	err := Parse(actual)
	So(err, ShouldBeNil)
	So(actual, ShouldResemble, expected)
}

func getRandFloat(size int) float64 {
	switch size {
	case 32:
		return rand.Float64()
	case 64:
		return float64(rand.Float32())
	default:
		panic("invalid size")
	}
}
