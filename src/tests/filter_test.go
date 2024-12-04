package tests

import (
	"github.com/onursedef/linqed-go/src"
	"testing"
)

func TestWhere(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Where(func(i int) bool {
		return i > 3
	}).Iterate()
	expected := []int{4, 5}

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}

	result = query.Where(func(i int) bool {
		return i > 10
	}).Iterate()

	if len(result) != 0 {
		t.Errorf("Expected [], got %v", result)
	}
}

func TestAny(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	query := linqed_go.From(items)

	result := query.Where(func(i int) bool {
		return i > 3
	}).Any()
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	result = query.Where(func(i int) bool {
		return i > 10
	}).Any()

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestAll(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	query := linqed_go.From(items)

	result := query.Where(func(i int) bool {
		return i > 0
	}).All(func(i int) bool {
		return i > 0
	})
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	result = query.Where(func(i int) bool {
		return i > 3
	}).All(func(i int) bool {
		return i < 3
	})

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
