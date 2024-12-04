package tests

import (
	"github.com/onursedef/linqed-go/src"
	"reflect"
	"testing"
)

func TestOrderBy(t *testing.T) {
	items := []int{5, 3, 1, 4, 2}
	query := linqed_go.From(items)

	result := query.OrderBy(func(i, j int) bool { return i < j }).Iterate()
	expected := []any{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
