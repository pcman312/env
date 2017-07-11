package env

import (
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type TestConfig struct {
	MyBool bool `env:"mybool"`

	MyString string `env:"mystring"`

	MyInt      int           `env:"myint"`
	MyInt8     int8          `env:"myint8"`
	MyInt16    int16         `env:"myint16"`
	MyInt32    int32         `env:"myint32"`
	MyInt64    int64         `env:"myint64"`
	MyDuration time.Duration `env:"myduration"`

	MyUint   uint   `env:"myuint"`
	MyUint8  uint8  `env:"myuint8"`
	MyUint16 uint16 `env:"myuint16"`
	MyUint32 uint32 `env:"myuint32"`
	MyUint64 uint64 `env:"myuint64"`

	MyFloat32 float32 `env:"myfloat32"`
	MyFloat64 float64 `env:"myfloat64"`

	MyURL *url.URL `env:"myurl"`

	////////////////////////////////////
	// Slices
	MyBoolArr []bool   `env:"myboolarr"`
	MyStrArr  []string `env:"mystrarr"`

	MyIntArr      []int           `env:"myintarr"`
	MyInt8Arr     []int8          `env:"myint8arr"`
	MyInt16Arr    []int16         `env:"myint16arr"`
	MyInt32Arr    []int32         `env:"myint32arr"`
	MyInt64Arr    []int64         `env:"myint64arr"`
	MyDurationArr []time.Duration `env:"mydurationarr"`

	MyUintArr   []uint   `env:"myuintarr"`
	MyUint8Arr  []uint8  `env:"myuint8arr"`
	MyUint16Arr []uint16 `env:"myuint16arr"`
	MyUint32Arr []uint32 `env:"myuint32arr"`
	MyUint64Arr []uint64 `env:"myuint64arr"`

	MyFloat32Arr []float32 `env:"myfloat32arr"`
	MyFloat64Arr []float64 `env:"myfloat64arr"`

	MyURLArr []*url.URL `env:"myurlarr"`
}

func TestParse(t *testing.T) {
	Convey("Not a pointer", t, func() {
		tc := TestConfig{}
		err := Parse(tc)
		So(err, ShouldEqual, ErrNotStructPointer)
	})

	Convey("Points to a non-struct", t, func() {
		s := ""
		err := Parse(&s)
		So(err, ShouldEqual, ErrNotStructPointer)
	})

	Convey("Bad required tag", t, func() {
		type BadStruct struct {
			BadRequiredField string `env:"badfield" required:"asdf"`
		}

		actual := &BadStruct{}
		err := Parse(actual)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "either true or false")
	})

	Convey("Unsupported type", t, func() {
		type MyStruct struct {
			Str string
		}
		type BadStruct struct {
			UnsupportedType MyStruct `env:"unsupportedtype"`
		}

		actual := &BadStruct{}
		err := Parse(actual)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unsupported type")
	})

	Convey("Unsupported slice", t, func() {
		type MyStruct struct {
			Str string
		}
		type BadStruct struct {
			UnsupportedSlice []MyStruct `env:"unsupportedslice"`
		}

		actual := &BadStruct{}
		err := Parse(actual)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unsupported slice type")
	})

	Convey("Unsupported pointer", t, func() {
		type BadStruct struct {
			UnsupportedPointer *string `env:"unsupportedptr"`
		}

		actual := &BadStruct{}
		err := Parse(actual)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "unsupported pointer type")
	})

	Convey("Missing required field", t, func() {
		type BadStruct struct {
			RequiredField string `env:"reqField" required:"true"`
		}

		actual := &BadStruct{}
		err := Parse(actual)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "missing required variable")
	})
}

func testParseEquality(actual, expected interface{}, pass bool, env []string) {
	defer resetEnv(env)

	err := Parse(actual)
	if !pass {
		So(err, ShouldNotBeNil)
	} else {
		So(actual, ShouldResemble, expected)
	}
}

func resetEnv(environ []string) {
	os.Clearenv()
	for _, env := range environ {
		keyVal := strings.SplitN(env, "=", 2)
		os.Setenv(keyVal[0], keyVal[1])
	}
}
