package bisect

import (
	"strings"
)

// Comparator defines a function type for comparing two elements.
// It should return a negative value if a < b, 0 if a == b, and a positive value if a > b.
type Comparator[T any] func(a, b T) int

// InsortRight inserts an item x into a sorted slice a, keeping it sorted.
func InsortRight[T any](a *[]T, x T, lo, hi int, compare Comparator[T]) {
	slice := *a
	lo = BisectRight(*a, x, lo, hi, compare)

	*a = append(slice[:lo], append([]T{x}, slice[lo:]...)...)
}

// BisectRight returns the index where to insert item x in a sorted slice a,
// assuming the slice is sorted. The returned index ensures that all elements
// in a[:i] are <= x, and all elements in a[i:] are > x.
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

// InsortLeft inserts an item x into a sorted slice a, keeping it sorted.
// If x is already in a, it inserts it before the leftmost x.
func InsortLeft[T any](a *[]T, x T, lo, hi int, compare Comparator[T]) {
	slice := *a
	lo = BisectLeft(*a, x, lo, hi, compare)                   // Find the insertion index
	*a = append(slice[:lo], append([]T{x}, slice[lo:]...)...) // Insert at the correct position
}

// BisectLeft returns the index where to insert item x in sorted slice a,
// assuming a is sorted. All elements in a[:i] are < x, and all elements in a[i:] are >= x.
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

// Global aliases using wrapper functions
func Bisect[T any](a []T, x T, lo int, hi int, compare Comparator[T]) int {
	return BisectRight(a, x, lo, hi, compare)
}

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
