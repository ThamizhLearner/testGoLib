// Note: test code container file names ends with _test.go
//		Also, the test function name starts with Test

package greetings_test

import (
	"testing"

	greetings "github.com/ThamizhLearner/testGoLib"
)

func TestFunction(t *testing.T) {
	name := "Gladys"
	msg := greetings.Hello(name)
	if msg != "Hi, Gladys. Welcome!" {
		t.Error("Failed")
	}
}
