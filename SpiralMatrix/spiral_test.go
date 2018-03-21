package SpiralMatrix

import "testing"

func TestSpiral2(t *testing.T) {
	matrix := spiral(2)

	if len(matrix) != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, len(matrix))
	}
}

func TestSpiral3(t *testing.T) {
	matrix := spiral(3)

	if len(matrix) != 3 {
		t.Errorf("Expected length to be %v, but got %v", 3, len(matrix))
	}
}

func TestSpiral4(t *testing.T) {
	matrix := spiral(4)

	if len(matrix) != 4 {
		t.Errorf("Expected length to be %v, but got %v", 4, len(matrix))
	}
}
