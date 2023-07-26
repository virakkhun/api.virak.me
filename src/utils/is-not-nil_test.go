package utils

import (
	"errors"
	"testing"
)

func TestNotNil(t *testing.T) {
	err := errors.New("Error")

	if !IsNotNil(err) {
		t.Fatal(`case: not nil failed`)
	}
}

func TestCaseNil(t *testing.T) {
	var err error = nil

	if IsNotNil(err) {
		t.Fatal(`case: nil failed`)
	}
}
