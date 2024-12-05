package tests

import (
	"github.com/onursedef/linqed-go"
	"testing"
)

type person struct {
	Name string
	Age  int
}

func TestSelect(t *testing.T) {
	// Sample data: a list of Person objects
	people := []person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Jack", Age: 35},
	}

	// Use the Where method to filter people who are 30 or older
	query := linqed_go.From(people).Where(func(p person) bool {
		return p.Age >= 30
	})

	// Use the Select method to project Name field after filtering
	result := linqed_go.Select(query, func(p person) string {
		return p.Name
	}).Iterate()

	// Expected result: a list of names of people who are 30 or older
	expected := []string{"John", "Jack"}

	// Check if the lengths are equal
	if len(result) != len(expected) {
		t.Errorf("Expected length %v, got %v", len(expected), len(result))
	}

	// Check if each result matches the expected value
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}
