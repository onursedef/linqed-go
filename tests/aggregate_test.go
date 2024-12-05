package tests

import (
	"github.com/onursedef/linqed-go"
	"testing"
)

func TestCount(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Count()
	expected := 5

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSum(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Sum()
	expected := 15.0

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMinMax(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	_min := query.Min()
	_max := query.Max()

	if _min != 1 {
		t.Errorf("Expected 1, got %v", _min)
	}

	if _max != 5 {
		t.Errorf("Expected 5, got %v", _max)
	}
}
