package tests

import (
	"github.com/onursedef/linqed-go"
	"reflect"
	"testing"
)

func TestSkip(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Skip(2).Iterate()

	if len(result) != 3 {
		t.Errorf("Expected 3 items, got %d", len(result))
	}

	expected := []any{3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTake(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	query := linqed_go.From(items)

	result := query.Take(2).Iterate()

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	expected := []any{1, 2}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
