package predicates

import (
	"github.com/luckyluke66/FuncSlice/slices"
)

// Elem checks if an element is present in the slice.
// It returns true if the element is found, otherwise false.
//
// Example:
//   slice := []int{1, 2, 3, 4}
//   output := Elem(3, slice) // output will be true
func Elem[T comparable](elem T, slice []T) bool {
    for _, v := range slice {
        if v == elem {
            return true
        } 
    }
    return false
}

// NotElem checks if an element is not present in the slice.
// It returns true if the element is not found, otherwise false.
//
// Example:
//   slice := []int{1, 2, 3, 4}
//   output := NotElem(5, slice) // output will be true
func NotElem[T comparable](elem T, slice []T) bool {
    return !Elem(elem, slice)
}

// IsPrefixOf checks if the prefix slice is a prefix of the main slice.
// It returns true if the prefix is found at the beginning of the main slice, otherwise false.
//
// Example:
//   prefix := []int{1, 2}
//   slice := []int{1, 2, 3, 4}
//   output := IsPrefixOf(prefix, slice) // output will be true
func IsPrefixOf[T comparable](prefix []T, slice []T) bool {
    for i, v := range prefix {
        if v != slice[i] {
            return false
        }
    }
    return true
}

// IsSuffixOf checks if the suffix slice is a suffix of the main slice.
// It returns true if the suffix is found at the end of the main slice, otherwise false.
//
// Example:
//   suffix := []int{3, 4}
//   slice := []int{1, 2, 3, 4}
//   output := IsSuffixOf(suffix, slice) // output will be true

func IsSuffixOf[T comparable](suffix []T, slice []T) bool {
    return IsPrefixOf(slices.Reverse(suffix), slices.Reverse(slice))
}

// Any returns true if the predicate function returns true for any element in the slice.
// It returns false if the predicate function returns false for all elements.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Any(input, func(x int) bool { return x > 3 }) // output will be true
func Any[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	
	return false
}

// All returns true if the predicate function returns true for all elements in the slice.
// It returns false if the predicate function returns false for any element.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := All(input, func(x int) bool { return x > 0 }) // output will be true
func All[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}