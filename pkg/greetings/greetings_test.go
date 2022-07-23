// SPDX-License-Identifier: AGPL-3.0-or-later
package greetings

import (
	"regexp"
	"testing"
)

func TestHelloNameNotNull(t *testing.T) {
	name := "liming"

	message, err := Hello(name)

	if err != nil {
		t.Log("not empty input should not throw error")
		t.Fail()
	}
	if len(message) == 0 {
		t.Log("not empty input should return not-nil value")
		t.Fail()
	}
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(message) {
		t.Fatalf("Input Value should Return")
	}
}

func TestHelloNameEmpty(t *testing.T) {
	name := ""

	message, err := Hello(name)

	if err == nil {
		t.Log("empty input should throw error")
		t.Fail()
	}
	if len(message) != 0 {
		t.Log("empty input should return empty value")
		t.Fail()
	}
}
