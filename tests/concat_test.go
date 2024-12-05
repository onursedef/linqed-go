package tests

import (
	"github.com/onursedef/linqed-go"
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	items1 := []int{1, 2, 3}
	items2 := []int{4, 5, 6}
	query := linqed_go.From(items1)

	result := query.Concat(items2).Iterate()
	expected := []any{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
