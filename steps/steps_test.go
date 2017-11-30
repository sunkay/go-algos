package steps

import "testing"

func TestStep2(t *testing.T) {
	steps := steps(2)
	if len(steps) != 2 {
		t.Errorf("Expected length to be %v, but got %v", 2, len(steps))
	}
	if steps[0] != "# " {
		t.Errorf("Expected step to be %v, but got %v", "# ", steps[0])
	}
	if steps[1] != "##" {
		t.Errorf("Expected step to be %v, but got %v", "##", steps[1])
	}
}
func TestStep3(t *testing.T) {
	steps := steps(3)
	if len(steps) != 3 {
		t.Errorf("Expected length to be %v, but got %v", 3, len(steps))
	}
	if steps[0] != "#  " {
		t.Errorf("Expected step to be %v, but got %v", "#  ", steps[0])
	}
	if steps[1] != "## " {
		t.Errorf("Expected step to be %v, but got %v", "## ", steps[1])
	}
	if steps[2] != "###" {
		t.Errorf("Expected step to be %v, but got %v", "###", steps[2])
	}
}
