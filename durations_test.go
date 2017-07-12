package env

import (
	"os"
	"testing"
	"time"

	"strings"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_duration(t *testing.T) {
	Convey("positive duration", t, func() {
		defer resetEnv(os.Environ())

		rawDur := "1h"
		dur, err := time.ParseDuration(rawDur)
		So(err, ShouldBeNil)
		actual := &TestConfig{}
		expected := &TestConfig{
			MyDuration: dur,
		}

		os.Setenv("myduration", rawDur)
		err = Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("duration slice", t, func() {
		defer resetEnv(os.Environ())

		rawDurs := []string{
			"1ns",
			"1us",
			"1ms",
			"1s",
			"1m",
			"1h",
		}
		durs := make([]time.Duration, len(rawDurs), len(rawDurs))
		for i, rawDur := range rawDurs {
			dur, err := time.ParseDuration(rawDur)
			So(err, ShouldBeNil)
			durs[i] = dur
		}
		actual := &TestConfig{}
		expected := &TestConfig{
			MyDurationArr: durs,
		}

		os.Setenv("mydurationarr", strings.Join(rawDurs, ", "))
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("duration min/max", t, func() {
		type TestStruct struct {
			MyDur time.Duration `env:"mydur" min:"1m" max:"1h"`
		}

		tests := map[string]bool{
			"1s":       false,
			"59s":      false,
			"59.9s":    false,
			"1m":       true,
			"1m01.1s":  true,
			"5m":       true,
			"30m":      true,
			"59m":      true,
			"59m59.9s": true,
			"60m":      true,
			"60m0.01s": false,
			"60m1s":    false,
			"61m":      false,
			"90m":      false,
		}

		for value, pass := range tests {
			actual := &TestStruct{}

			dur, err := time.ParseDuration(value)
			So(err, ShouldBeNil)
			expected := &TestStruct{
				MyDur: dur,
			}
			env := os.Environ()
			os.Setenv("mydur", fmt.Sprint(value))
			testParseEquality(actual, expected, pass, env)
		}
	})
}
