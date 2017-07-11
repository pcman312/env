package env

import (
	"net/url"
	"reflect"

	"github.com/pcman312/errutils"
)

func handleUrl(ref reflect.Value, rawVal string) error {
	if rawVal == "" {
		return nil
	}
	u, err := url.Parse(rawVal)
	if err != nil {
		return err
	}
	ref.Set(reflect.ValueOf(u))
	return nil
}

func handleUrlSlice(ref reflect.Value, rawArr []string) error {
	if len(rawArr) == 0 {
		return nil
	}
	urls := make([]*url.URL, 0, len(rawArr))
	errs := []error{}

	for _, str := range rawArr {
		url, err := url.Parse(str)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		urls = append(urls, url)
	}

	if len(errs) > 0 {
		return errutils.JoinErrs(", ", errs...)
	}

	ref.Set(reflect.ValueOf(urls))
	return nil
}
