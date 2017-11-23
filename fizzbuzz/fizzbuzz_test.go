package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz15(t *testing.T) {
	r := fizzBuzz(10)

	if r[0] != 1 {
		t.Errorf("Expected r[0] to be 1 but got %v", r[0])
	}
	if r[1] != 2 {
		t.Errorf("Expected r[1] to be 1 but got %v", r[1])
	}
	if r[2] != "fizz" {
		t.Errorf("Expected r[2] to be fizz but got %v", r[2])
	}
	if r[3] != 4 {
		t.Errorf("Expected r[3] to be 4 but got %v", r[3])
	}
	if r[4] != "buzz" {
		t.Errorf("Expected r[4] to be bizz but got %v", r[4])
	}
	if r[5] != "fizz" {
		t.Errorf("Expected r[5] to be fizz but got %v", r[5])
	}
	if r[6] != 7 {
		t.Errorf("Expected r[6] to be 1 but got %v", r[6])
	}
	if r[7] != 8 {
		t.Errorf("Expected r[7] to be 1 but got %v", r[7])
	}
	if r[8] != "fizz" {
		t.Errorf("Expected r[8] to be bizz but got %v", r[8])
	}
	if r[9] != "buzz" {
		t.Errorf("Expected r[9] to be fizz but got %v", r[9])
	}

}
