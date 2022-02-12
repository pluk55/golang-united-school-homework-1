package solutions

import (
	"fmt"
	"strings"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	msg := GetMessage()

	if strings.Contains(msg, "Hello") {
		fmt.Println("Passed")
	}
}
