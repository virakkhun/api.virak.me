package utils

import "testing"

func TestCaseEmpty(t *testing.T) {
	emptyString := ""

	if !IsStringEmpty(emptyString) {
		t.Fatal("case: empty string failed")
	}
}

func TestCaseNotEmpty(t *testing.T) {
	str := "Hello"

	if IsStringEmpty(str) {
		t.Fatalf("case: not empty string failed")
	}
}
