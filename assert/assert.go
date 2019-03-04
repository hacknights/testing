package assert

import goTesting "testing"

// Error unwraps the message in error and calls Fail
func Error(t *goTesting.T, expected interface{}, err error) {
	Fail(t, expected, err.Error())
}

// Fail wraps the call to t.Errorf and provides
// a consistent message for failed tests by
// accepting the expected and actual values
func Fail(t *goTesting.T, expected, actual interface{}) {
	t.Errorf("expected: %s, actual %s", expected, actual)
}

// Log wraps the call to t.Log
func Log(t *goTesting.T, args ...interface{}) {
	t.Log(args...)
}
