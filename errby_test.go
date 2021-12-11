package errby_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/matryer/is"

	"github.com/elblox/errby"
	"github.com/elblox/errby/mock"
)

func TestErrBy(t *testing.T) {
	tt := []struct {
		actualErr, expectedErr            error
		fail, failContains, failMustMatch bool
	}{
		{
			actualErr:     fmt.Errorf("foo"),
			expectedErr:   fmt.Errorf("bar"),
			fail:          true,
			failContains:  true,
			failMustMatch: true,
		},
		{
			actualErr:     fmt.Errorf("foo"),
			expectedErr:   fmt.Errorf("foo"),
			fail:          false,
			failContains:  false,
			failMustMatch: false,
		},
		{
			actualErr:     nil,
			expectedErr:   nil,
			fail:          false,
			failContains:  false,
			failMustMatch: false,
		},
		{
			actualErr:     nil,
			expectedErr:   fmt.Errorf("foo"),
			fail:          true,
			failContains:  true,
			failMustMatch: true,
		},
		{
			actualErr:     fmt.Errorf("foo"),
			expectedErr:   nil,
			fail:          true,
			failContains:  true,
			failMustMatch: true,
		},
		{
			actualErr:     fmt.Errorf("an error, root cause: foo"),
			expectedErr:   fmt.Errorf("foo"),
			fail:          true,
			failContains:  false,
			failMustMatch: false,
		},
		{
			actualErr:     fmt.Errorf("an error, root cause: foo"),
			expectedErr:   fmt.Errorf("an [a-z]+, root cause: foo"),
			fail:          true,
			failContains:  true,
			failMustMatch: false,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Compare %v vs. %v", tc.expectedErr, tc.actualErr), func(t *testing.T) {
			tmock := &mock.TMock{FailNowFunc: func() {}}

			ismock, deferfunc, err := silentIs(t, tmock)
			if err != nil {
				t.Fatalf("failed to create silent is: %v", err)
			}
			defer deferfunc()

			is := is.New(t)

			errby.Compare(ismock, tc.expectedErr, tc.actualErr)
			is.Equal(tc.fail, len(tmock.FailNowCalls()) > 0)
		})

		t.Run(fmt.Sprintf("Contains %v vs. %v", tc.expectedErr, tc.actualErr), func(t *testing.T) {
			tmock := &mock.TMock{FailNowFunc: func() {}}

			ismock, deferfunc, err := silentIs(t, tmock)
			if err != nil {
				t.Fatalf("failed to create silent is: %v", err)
			}
			defer deferfunc()

			is := is.New(t)

			errby.Contains(ismock, tc.expectedErr, tc.actualErr)
			is.Equal(tc.failContains, len(tmock.FailNowCalls()) > 0)
		})

		t.Run(fmt.Sprintf("MustMatch %v vs. %v", tc.expectedErr, tc.actualErr), func(t *testing.T) {
			tmock := &mock.TMock{FailNowFunc: func() {}}

			ismock, deferfunc, err := silentIs(t, tmock)
			if err != nil {
				t.Fatalf("failed to create silent is: %v", err)
			}
			defer deferfunc()

			is := is.New(t)

			errby.MustMatch(ismock, tc.expectedErr, tc.actualErr)
			is.Equal(tc.failMustMatch, len(tmock.FailNowCalls()) > 0)
		})
	}
}

func silentIs(t *testing.T, isT is.T) (*is.I, func(), error) {
	t.Helper()

	out, err := ioutil.TempFile("./", "silentisout_")
	if err != nil {
		return &is.I{}, nil, fmt.Errorf("failed to create temp file for ismock redirect: %v", err)
	}
	deferfunc := func() {
		defererr := out.Close()
		if defererr != nil {
			t.Logf("failed to close temp file: %v", defererr)
		}
		defererr = os.Remove(out.Name())
		if defererr != nil {
			t.Logf("failed to delete temp file: %v", defererr)
		}
	}
	stdoutBak := os.Stdout
	os.Stdout = out
	ismock := is.New(isT)
	os.Stdout = stdoutBak

	return ismock, deferfunc, nil
}
