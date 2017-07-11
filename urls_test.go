package env

import (
	"os"
	"strings"
	"testing"

	"net/url"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse_urls(t *testing.T) {
	Convey("url", t, func() {
		defer resetEnv(os.Environ())

		url, err := url.Parse("http://www.google.com")
		So(err, ShouldBeNil)
		expected := &TestConfig{
			MyURL: url,
		}
		actual := &TestConfig{}

		os.Setenv("myurl", url.String())
		err = Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("url slice", t, func() {
		defer resetEnv(os.Environ())

		rawUrls := []string{
			"http://www.google.com",
			"http://www.reddit.com",
			"192.168.0.1",
		}
		urls := make([]*url.URL, len(rawUrls), len(rawUrls))

		for i, rawUrl := range rawUrls {
			url, err := url.Parse(rawUrl)
			So(err, ShouldBeNil)
			urls[i] = url
		}
		expected := &TestConfig{
			MyURLArr: urls,
		}
		actual := &TestConfig{}

		os.Setenv("myurlarr", strings.Join(rawUrls, ", "))
		err := Parse(actual)
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})
}
