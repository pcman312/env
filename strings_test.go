package env

import (
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_string(t *testing.T) {
	Convey("String", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyString: "omgwtfbbq",
		}

		os.Setenv("mystring", expected.MyString)
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("String slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyStrArr: []string{"one", "two", "three"},
		}

		os.Setenv("mystrarr", strings.Join(expected.MyStrArr, ", "))
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})
}
