package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	emg := "beer"
	msg := GetMessage(emg)

	if strings.Contains(msg, emg) {
		fmt.Println("Passed")
	}
}

func TestHelloEmpty(t *testing.T) {
	emg := "stamm"
	msg := GetMessage(emg)

	if strings.Contains(msg, emg) {
		t.Fatal()
	}
}
