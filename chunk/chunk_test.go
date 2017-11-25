package chunk

import (
	"reflect"
	"testing"
)

func TestChunk4Size2(t *testing.T) {
	compareSlices(chunk([]int{1, 2, 3, 4}, 2), [][]int{{1, 2}, {3, 4}}, t)
}
func TestChunk10Size2(t *testing.T) {
	compareSlices(chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2), [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}, t)
}

func TestChunk3Size1(t *testing.T) {
	compareSlices(chunk([]int{1, 2, 3}, 1), [][]int{{1}, {2}, {3}}, t)
}

func TestChunk5Size3(t *testing.T) {
	compareSlices(chunk([]int{1, 2, 3, 4, 5}, 3), [][]int{{1, 2, 3}, {4, 5}}, t)
}

func TestChunk13Size5(t *testing.T) {
	compareSlices(chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 5),
		[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}}, t)
}

func compareSlices(a [][]int, b [][]int, t *testing.T) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v, but got %v", b, a)
	}
}
