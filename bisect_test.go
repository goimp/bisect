package bisect

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func TestBisectRightInt(t *testing.T) {
	a := []int{1, 2, 4, 4, 5}
	x := 4

	// Find insertion index
	index := Bisect(a, x, 0, -1, CompareInt)
	if index != 4 {
		t.Errorf("Expected 4, found %d", index)
	}

	// Find insertion index for a different value
	index = Bisect(a, 3, 0, -1, CompareInt)
	if index != 2 {
		t.Errorf("Expected 2, found %d", index)
	}
}

func TestInsortRightInt(t *testing.T) {
	a := []int{1, 2, 4, 4, 5}
	x := 4

	// Insert into the slice
	Insort(&a, x, 0, -1, CompareInt)
	expected := []int{1, 2, 4, 4, 4, 5}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("expected %v, got %v", expected, a)
	}

	// Insert another value
	Insort(&a, 3, 0, -1, CompareInt)
	expected = []int{1, 2, 3, 4, 4, 4, 5}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("expected %v, got %v", expected, a)
	}
}

func TestInsortRightString(t *testing.T) {
	// Slice of strings sorted lexicographically
	words := []string{"apple", "banana", "cherry", "date"}
	expected := []string{"apple", "banana", "blueberry", "cherry", "date"}
	// New word to insert
	newWord := "blueberry"

	// Insert the new word
	Insort(&words, newWord, 0, -1, CompareStrings)

	if !reflect.DeepEqual(words, expected) {
		t.Errorf("expected %v, got %v", expected, words)
	}

}

func TestInsortRightStructByKey(t *testing.T) {
	// Slice of people sorted by age
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	expected := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"David", 30},
		{"Charlie", 35},
	}

	// New person to insert
	newPerson := Person{"David", 30}

	// Comparator for sorting by age
	compareByAge := func(a, b Person) int {
		return a.Age - b.Age
	}

	// Insert the new person
	Insort(&people, newPerson, 0, -1, compareByAge)

	if !reflect.DeepEqual(people, expected) {
		t.Errorf("expected %v, got %v", expected, people)
	}
}

func TestBisectLeftInt(t *testing.T) {
	a := []int{1, 2, 4, 4, 5}
	x := 4

	// Find insertion index
	index := BisectLeft(a, x, 0, -1, CompareInt)
	if index != 2 {
		t.Errorf("Expected 2, found %d", index)
	}

	// Find insertion index for a different value
	index = BisectLeft(a, 3, 0, -1, CompareInt)
	if index != 2 {
		t.Errorf("Expected 2, found %d", index)
	}
}

func TestInsortLeftInt(t *testing.T) {
	a := []int{1, 2, 4, 4, 5}
	x := 4

	// Insert into the slice
	InsortLeft(&a, x, 0, -1, CompareInt)
	expected := []int{1, 2, 4, 4, 4, 5}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("expected %v, got %v", expected, a)
	}

	// Insert another value
	InsortLeft(&a, 3, 0, -1, CompareInt)
	expected = []int{1, 2, 3, 4, 4, 4, 5}
	if !reflect.DeepEqual(a, expected) {
		t.Errorf("expected %v, got %v", expected, a)
	}
}

func TestInsortLeftSliceOfSlicesExample(t *testing.T) {

	// Comparator to compare slices by their first element (or any other criteria)
	compareByFirstElement := func(a, b []int) int {
		// Comparing the first element of each slice (list of ints)
		return a[0] - b[0]
	}

	// Comparator to compare slices of strings lexicographically (first element comparison)
	compareByFirstString := func(a, b []string) int {
		return strings.Compare(a[0], b[0])
	}

	// Slice of lists of integers (a list of lists)
	a := [][]int{
		{3, 4, 5},
		{1, 2, 3},
		{5, 6, 7},
	}

	expected_a := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{5, 6, 7},
	}

	// Sort by the first element of each list (slice of ints)
	sort.Slice(a, func(i, j int) bool {
		return compareByFirstElement(a[i], a[j]) < 0
	})
	fmt.Println("Sorted by first element of each sublist:")
	for _, sublist := range a {
		fmt.Println(sublist)
	}

	// Insert a new sublist into the sorted slice of lists
	newSublist := []int{2, 3, 4}
	InsortLeft(&a, newSublist, 0, -1, compareByFirstElement)
	fmt.Println("\nAfter Inserting new sublist (sorted by first element):")
	for _, sublist := range a {
		fmt.Println(sublist)
	}

	if !reflect.DeepEqual(a, expected_a) {
		t.Errorf("expected %v, got %v", expected_a, a)
	}

	// Slice of lists of strings (a list of lists of strings)
	b := [][]string{
		{"apple", "banana"},
		{"kiwi", "orange"},
		{"cherry", "date"},
	}

	expected_b := [][]string{
		{"apple", "banana"},
		{"cherry", "date"},
		{"grape", "melon"},
		{"kiwi", "orange"},
	}

	// Sort by the first string in each list (lexicographically)
	sort.Slice(b, func(i, j int) bool {
		return compareByFirstString(b[i], b[j]) < 0
	})
	fmt.Println("\nSorted by first string of each sublist:")
	for _, sublist := range b {
		fmt.Println(sublist)
	}

	// Insert a new sublist of strings
	newStringSublist := []string{"grape", "melon"}
	InsortLeft(&b, newStringSublist, 0, -1, compareByFirstString)
	fmt.Println("\nAfter Inserting new sublist of strings (sorted by first string):")
	for _, sublist := range b {
		fmt.Println(sublist)
	}

	if !reflect.DeepEqual(b, expected_b) {
		t.Errorf("expected %v, got %v", expected_b, b)
	}
}
