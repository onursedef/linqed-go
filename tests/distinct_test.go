package tests

import (
	"github.com/onursedef/linqed-go"
	"testing"
)

func TestDistinct(t *testing.T) {
	items := []int{1, 2, 3, 2, 4, 1, 5}
	query := linqed_go.From(items)

	result := query.Distinct().Iterate()
	expected := []int{1, 2, 3, 4, 5}

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
