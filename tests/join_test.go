package tests

import (
	"github.com/onursedef/linqed-go"
	"testing"
)

type Person struct {
	Name string
	City string
}

type City struct {
	Name       string
	Population int
}

func TestJoin(t *testing.T) {
	people := []Person{
		{Name: "John", City: "New York"},
		{Name: "Jane", City: "Los Angeles"},
		{Name: "Jack", City: "Chicago"},
	}

	cities := []City{
		{Name: "New York", Population: 10000000},
		{Name: "Los Angeles", Population: 4000000},
		{Name: "Chicago", Population: 2700000},
	}

	query := linqed_go.From(people)
	result := linqed_go.Join(query, cities,
		func(p Person) string { return p.City },
		func(c City) string { return c.Name },
		func(p Person, c City) struct {
			Name       string
			City       string
			Population int
		} {
			return struct {
				Name       string
				City       string
				Population int
			}{p.Name, c.Name, c.Population}
		}).Iterate()

	expected := []any{
		struct {
			Name       string
			City       string
			Population int
		}{
			Name: "John", City: "New York", Population: 10000000,
		},
		struct {
			Name       string
			City       string
			Population int
		}{
			Name: "Jane", City: "Los Angeles", Population: 4000000,
		},
		struct {
			Name       string
			City       string
			Population int
		}{
			Name: "Jack", City: "Chicago", Population: 2700000,
		},
	}

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}
