// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-15 01:02:33 A5E71B                         memd/[support_test.go]
// -----------------------------------------------------------------------------

package fixed64

// # Test Helpers
//   detectError(args ...interface{}) error
//   mockError(enable bool)
//   testErrors(label string, exec func() error, fail func(erm string))
//   testErrorsHelper(
//       label string,
//       wantError bool,
//       exec func() error,
//       fail func(erm string),
//   )
//   testGot(label string, got, want interface{}, fail func(erm string))
//   testLine() int
//   testNoErrors(label string, exec func() error, fail func(erm string))

// generate a test coverage report:
// go test -coverprofile cover.out
// go tool cover -html=cover.out

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// gotErm stores the last error message detacted by detectError()
var gotErm string

// gotError is a flag used for testing if errors are detected
var gotError bool

// NoError is a placeholder error that is returned in
// the closure function called by testErrors().
// If the function just returned nil, testErrors() would
// complain no error was returned. But not all functions
// need to return an error (even if they detect
// inconsistencies). In these cases just return NoError.
var NoError = errors.New("No Error")

// -----------------------------------------------------------------------------
// # Test Helpers

// detectError sets gotError to true when called
// (this function is only called by mockError() to test error trapping)
func detectError(args ...interface{}) error {
	gotError = true
	gotErm = fmt.Sprint(args...)
	return errors.New(gotErm)
}

// mockError changes the module's error handler to detectError() so unit
// tests can check if various error conditions are detected properly.
func mockError(enable bool) {
	gotError = false
	gotErm = ""
	if enable {
		mod.Error = detectError
		return
	}
	mod.Error = logError
}

// testErrors tests if a block of code specified in exec() function
// detects errors by calling mod.Error() and returns an error value.
//
// If a function doesn't return an error value,
// then exec() should return NoError.
//
// If the test fails, calls fail() function passing it
// the label combined with the error message string.
//
func testErrors(label string, exec func() error, fail func(erm string)) {
	testErrorsHelper(label, true, exec, fail)
}

// testErrorsHelper does the work for testErrors() and testNoErrors()
func testErrorsHelper(
	label string,
	wantError bool,
	exec func() error,
	fail func(erm string),
) {
	mockError(true)
	var (
		err = exec()
	)
	var erm string
	switch {
	case wantError && !gotError:
		{
			erm = "failed to detect error"
		}
	case wantError && err == nil:
		{
			erm = "failed to return error"
		}
	case !wantError && gotError:
		{
			erm = "detected unexpected error:\n"
			if len(gotErm) == 0 {
				erm += "<nil>"
			} else {
				erm += gotErm
			}
		}
	case !wantError && err != nil:
		{
			erm = "returned unexpected error:\n"
			if err == nil {
				erm += "<nil>"
			} else {
				erm += err.Error()
			}
		}
	}
	if len(erm) > 0 {
		erm = label + " " + erm
		fail(erm)
	}
	mockError(false)
}

// testGot compares 'got' to 'want' and calls fail() if they don't match.
//
// Call testGot like this, after you've prepared label, got and want vars:
// testGot(label, got, want, func(erm string) { t.Error(erm) })
//
func testGot(label string, got, want interface{}, fail func(erm string)) {
	var (
		gotS  = fmt.Sprint(got)
		wantS = fmt.Sprint(want)
	)
	if gotS == wantS {
		return
	}
	if strings.HasSuffix(label, "=") {
		erm := fmt.Sprintf("%s `%s`. want `%s`", label, gotS, wantS)
		fail(erm)
		return
	}
	erm := fmt.Sprintf("%s failed:\n\n"+
		"got:\n"+"`%s`\n\n"+
		"want:\n"+"`%s`\n\n",
		label, gotS, wantS)
	fail(erm)
}

// testLine returns the line number in a top-level test
func testLine() int {
	var ret int
	for i := 1; ; i++ {
		var (
			pc, _, lineNo, _ = runtime.Caller(i)
			funcName         = runtime.FuncForPC(pc).Name()
		)
		if funcName == "" || strings.HasPrefix(funcName, "testing.") {
			break
		}
		ret = lineNo
	}
	return ret
}

// testNoErrors tests if a block of code specified in exec() function
// detects no errors (mod.Error() isn't called) and returns no error value.
//
// If a function doesn't return an error value,
// then exec() should return nil.
//
// If the test fails, calls fail() function passing it
// the label combined with the error message string.
//
func testNoErrors(label string, exec func() error, fail func(erm string)) {
	testErrorsHelper(label, false, exec, fail)
}

//end
