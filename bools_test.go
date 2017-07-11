package env

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_bool(t *testing.T) {
	Convey("Boolean", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyBool: true,
		}

		os.Setenv("mybool", "true")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("Boolean slice", t, func() {
		defer resetEnv(os.Environ())

		actual := &TestConfig{}
		expected := &TestConfig{
			MyBoolArr: []bool{true, false, true, true, false},
		}

		os.Setenv("myboolarr", "true, false, true, true, false")
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})
}
