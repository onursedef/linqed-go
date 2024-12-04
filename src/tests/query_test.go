package tests

import (
	"github.com/onursedef/linqed-go/src"
	"testing"
)

func TestFirst(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result, found := query.Where(func(i int) bool {
		return i > 3
	}).First()
	expected := 4

	if found != nil {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result, found = query.Where(func(i int) bool {
		return i > 10
	}).First()
	if found == nil {
		t.Errorf("Expected error, got %v", result)
	}
}

func TestFirstOrDefault(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Where(func(i int) bool {
		return i > 3
	}).FirstOrDefault()
	expected := 4

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result = query.Where(func(i int) bool {
		return i > 10
	}).FirstOrDefault()

	if result != 0 {
		t.Errorf("Expected 0, got %v", result)
	}
}
