package errby

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/matryer/is"
)

// Compare provides a convinient function to compare errors as
// helper in tests.
func Compare(is *is.I, expected, actual error) {
	is.Helper()
	if expected != nil && actual != nil {
		is.Equal(expected.Error(), actual.Error())
		return
	}
	is.Equal(expected, actual)
}

// Contains provides a convenient function to compare errors based on
// string.Contains as helper in tests.
func Contains(is *is.I, expected, actual error) {
	is.Helper()
	if expected != nil && actual != nil {
		if !strings.Contains(actual.Error(), expected.Error()) {
			is.NoErr(fmt.Errorf("%q is not contained in %q", expected.Error(), actual.Error()))
		}
		return
	}
	is.Equal(expected, actual)
}

// Match provides a convinient function to compare errors based on
// regexp.MatchString as helper in tests.
func MustMatch(is *is.I, expected, actual error) {
	is.Helper()
	if expected != nil && actual != nil {
		re := regexp.MustCompile(expected.Error())
		if !re.MatchString(actual.Error()) {
			is.NoErr(fmt.Errorf("%q is not contained in %q", expected.Error(), actual.Error()))
		}
		return
	}
	is.Equal(expected, actual)
}

//go:generate moq -fmt goimports -pkg mock -out mock/is_T.go . T

// T reports when failures occur.
// testing.T implements this interface.
type T interface {
	Fail()
	FailNow()
}
