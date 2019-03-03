package testing

import "testing"

// Fail wraps the call to t.Errorf and provides
// a consistent message for failed tests by
// accepting the expected and actual values
func Fail(t *testing.T, expected, actual interface{}) {
	t.Errorf("expected: %s, actual %s", expected, actual)
}

// Log wraps the call to t.Log
func Log(t *testing.T, args ...interface{}) {
	t.Log(args...)
}
