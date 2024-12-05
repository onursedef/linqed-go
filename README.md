# LINQed-Go
[![Go Reference](https://pkg.go.dev/badge/github.com/onursedef/linqed-go.svg)](https://pkg.go.dev/github.com/onursedef/linqed-go) [![Coverage Status](https://coveralls.io/repos/github/onursedef/linqed-go/badge.svg?branch=main)](https://coveralls.io/github/onursedef/linqed-go?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/onursedef/linqed-go)](https://goreportcard.com/report/github.com/onursedef/linqed-go)

`LINQed-Go` is a lightweight Go package that brings LINQ-like functionality to Go, allowing you to perform common operations like filtering, projecting, ordering, and joining on collections in a fluent, method-chaining style.

## Features

- **Where**: Filter a collection based on a condition (similar to LINQ's `Where`).
- **Select**: Transform a collection's items (similar to LINQ's `Select`).
- **OrderBy**: Sort a collection based on a provided key selector (similar to LINQ's `OrderBy`).
- **Skip**: Skip a specified number of elements from the start.
- **Take**: Take a specified number of elements from the start.
- **Join**: Perform an inner join between two collections based on a key.
- **First**: Get the first element of the collection.
- **FirstOrDefault**: Get the first element, or a default value if no elements match.
- **Any**: Check if any elements satisfy a condition.
- **All**: Check if all elements satisfy a condition.
- **Count**: Count the number of elements in the collection.
- **Distinct**: Get distinct elements from the collection.
- **Sum, Min, Max**: Aggregate functions for summing, getting the minimum, or maximum values.
- **Concat**: Concatenate two collections.

## Installation

To install the package, use the following Go command:

```bash
go get github.com/onursedef/linqed-go
```

## Usage

Here's a simple example of how you can use `LINQed-Go`:

```go
package main

import (
	"fmt"
	"github.com/onursedef/linqed-go"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Sample data
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Jack", Age: 35},
	}

	// Create a queryable collection
	query := linqed_go.From(people).Where(func(p Person) bool {
        return p.Age >= 30
    })

	// Example of Where and Select methods
	result := linqed_go.Select(query, func(p Person) string {
        return p.Name
    }).Iterate()

	fmt.Println(result) // Output: [John Jack]
}
```

## Fluent API
`LINQed-Go` uses a fluent API to chain methods together, enabling an expressive syntax for querying collections. The primary methods are:

- **From**: Initialize a new queryable collection from a slice.
- **Where**: Filter items based on a condition.
- **Select**: Project items into a new form.
- **OrderBy**: Sort items by a key.
- **Skip**: Skip the first N elements.
- **Take**: Take the first N elements.
- **Join**: Perform an inner join between two collections.
- **Iterate**: Execute the query and return the results.

## Methods

### From
```go
func From[T any](items []T) *Queryable[T]
```
Creates a new `Queryable` collection from a slice of items.

### Where
```go
func (q *Queryable[T]) Where(predicate func(T) bool) *Queryable[T]
```
Filters the collection based on a predicate function.

### Select
```go
func (q *Queryable[T]) Select(projection func(T) any) *Queryable[T]
```
Projects each item in the collection into a new form.
**Note:** This method is not chainable.

### OrderBy
```go
func (q *Queryable[T]) OrderBy(comparator func(i, j T) bool) *Queryable[T]
```
Orders the collection based on the provided key selector.

### Skip
```go
func (q *Queryable[T]) Skip(n int) *Queryable[T]
```
Skips the first `count` elements from the collection.

### Take
```go
func (q *Queryable[T]) Take(n int) *Queryable[T]
```
Takes the first `count` elements from the collection.

### Join
```go
func Join[T, U any, K comparable, R any](q *Queryable[T], other []U, keySelector func(T) K, otherKeySelector func(U) K, resultSelector func(T, U) R) *Queryable[R]
```
Performs an inner join between two collections based on a key.
**Note:** This method is not chainable.

### First
```go
func (q *Queryable[T]) First() (T, error)
```
Gets the first element of the collection.

### FirstOrDefault
```go
func (q *Queryable[T]) FirstOrDefault() T
```
Gets the first element of the collection, or a default value if no elements match.

### Any
```go
func (q *Queryable[T]) Any() bool
```
Checks if any elements satisfy a condition.

### All
```go
func (q *Queryable[T]) All(predicate func(T) bool) bool
```
Checks if all elements satisfy a condition.

### Count
```go
func (q *Queryable[T]) Count() int
```
Counts the number of elements in the collection.

### Distinct
```go
func (q *Queryable[T]) Distinct() *Queryable[T]
```
Gets distinct elements from the collection.

### Sum, Min, Max
```go
func (q *Queryable[T]) Sum() float64
func (q *Queryable[T]) Max() float64
func (q *Queryable[T]) Min() float64
```
Aggregate functions for summing, getting the minimum, or maximum values.

### Concat
```go
func (q *Queryable[T]) Concat(other []T) *Queryable[T]
```
Concatenates two collections.

### Iterate
```go
func (q *Queryable[T]) Iterate() []T
```
Executes the query and returns the results.

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request if you have any improvements or suggestions.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.