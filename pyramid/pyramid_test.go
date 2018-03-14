package pyramid

import (
	"testing"
)

func TestPyramid2(t *testing.T) {
	var steps []string
	steps = pyramid(2)

	if len(steps) != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, len(steps))
	}
	if steps[0] != " # " {
		t.Errorf("Expected step to be %v, but got %v", " # ", steps[0])
	}
	if steps[1] != "###" {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
}

func TestPyramid3(t *testing.T) {
	var steps []string
	steps = pyramid(3)

	if len(steps) != 3 {
		t.Errorf("Expected length to be %v, but got %v", 3, len(steps))
	}
	if steps[0] != "  #  " {
		t.Errorf("Expected step to be %v, but got %v", " # ", steps[0])
	}
	if steps[1] != " ### " {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
	if steps[2] != "#####" {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
}

func TestPyramid4(t *testing.T) {
	var steps []string
	steps = pyramid(4)

	if len(steps) != 4 {
		t.Errorf("Expected length to be %v, but got %v", 4, len(steps))
	}
	if steps[0] != "   #   " {
		t.Errorf("Expected step to be %v, but got %v", " # ", steps[0])
	}
	if steps[1] != "  ###  " {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
	if steps[2] != " ##### " {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
	if steps[3] != "#######" {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[1])
	}
}
