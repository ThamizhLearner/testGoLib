// Note: test code container file names ends with _test.go
//		Also, the test function name starts with Test

package greetings

import "testing"

func TestFunction(t *testing.T) {
	name := "Gladys"
	msg := Hello(name)
	if msg != "Hi, Gladys. Welcome!" {
		t.Error("Failed")
	}
}
