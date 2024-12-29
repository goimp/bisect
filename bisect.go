/*
Bisection algorithms
*/

package bisect

import (
	"strings"
)

// Comparator defines a function type for comparing two elements.
// It should return a negative value if a < b, 0 if a == b, and a positive value if a > b.
type Comparator[T any] func(a, b T) int

// Insert item x in list a, and keep it sorted assuming a is sorted.
// If x is already in a, insert it to the right of the rightmost x.
// Optional args lo (default 0) and hi (default len(a)) bound the
// slice of a to be searched.
// A comparator function should be supplied to customize the sort order.
func InsortRight[T any](a *[]T, x T, lo, hi int, compare Comparator[T]) {
	slice := *a
	lo = BisectRight(*a, x, lo, hi, compare)

	*a = append(slice[:lo], append([]T{x}, slice[lo:]...)...)
}

// Return the index where to insert item x in list a, assuming a is sorted.
// The return value i is such that all e in a[:i] have e <= x, and all e in
// a[i:] have e > x.  So if x already appears in the list, a.insert(i, x) will
// insert just after the rightmost x already there.
// Optional args lo (default 0) and hi (default len(a)) bound the
// slice of a to be searched.
// A comparator function should be supplied to customize the sort order.
func BisectRight[T any](a []T, x T, lo int, hi int, compare Comparator[T]) int {
	if lo < 0 {
		panic("lo must be non-negative")
	}
	if hi > len(a) {
		panic("hi cannot exceed slice length")
	}
	if hi == -1 {
		hi = len(a)
	}

	for lo < hi {
		mid := (lo + hi) / 2
		if compare(x, a[mid]) < 0 { // x < a[mid]
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

// Insert item x in list a, and keep it sorted assuming a is sorted.
// If x is already in a, insert it to the left of the leftmost x.
// Optional args lo (default 0) and hi (default len(a)) bound the
// slice of a to be searched.
// A comparator function should be supplied to customize the sort order.
func InsortLeft[T any](a *[]T, x T, lo, hi int, compare Comparator[T]) {
	slice := *a
	lo = BisectLeft(*a, x, lo, hi, compare)                   // Find the insertion index
	*a = append(slice[:lo], append([]T{x}, slice[lo:]...)...) // Insert at the correct position
}

// Return the index where to insert item x in list a, assuming a is sorted.
// The return value i is such that all e in a[:i] have e < x, and all e in
// a[i:] have e >= x.  So if x already appears in the list, a.insert(i, x) will
// insert just before the leftmost x already there.
// Optional args lo (default 0) and hi (default len(a)) bound the
// slice of a to be searched.
// A comparator function should be supplied to customize the sort order.
func BisectLeft[T any](a []T, x T, lo int, hi int, compare Comparator[T]) int {
	if lo < 0 {
		panic("lo must be non-negative")
	}
	if hi == -1 {
		hi = len(a)
	}

	for lo < hi {
		mid := (lo + hi) / 2
		if compare(a[mid], x) < 0 { // a[mid] < x
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

// Global alias to BisectRight using wrapper functions
func Bisect[T any](a []T, x T, lo int, hi int, compare Comparator[T]) int {
	return BisectRight(a, x, lo, hi, compare)
}

// Global alias to InsortRight using wrapper functions
func Insort[T any](a *[]T, x T, lo, hi int, compare Comparator[T]) {
	InsortRight(a, x, lo, hi, compare)
}

// Comparator for sorting strings lexicographically
func CompareStrings(a, b string) int {
	return strings.Compare(a, b)
}

// Comparator for integers
func CompareInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Comparator for unsigned integers
func CompareUint(a, b uint) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Comparator for floating-point numbers
func CompareFloat64(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Comparator for floating-point numbers
func CompareFloat32(a, b float32) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Comparator for boolean values (false < true)
func CompareBool(a, b bool) int {
	if a == b {
		return 0
	}
	if a {
		return 1
	}
	return -1
}

// Comparator for bytes (unsigned 8-bit integers)
func CompareByte(a, b byte) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Comparator for runes (Unicode code points)
func CompareRune(a, b rune) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
