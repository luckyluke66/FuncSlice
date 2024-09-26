package setOp

import (
    "github.com/luckyluke66/SliceLib/predicates"
)

// Nub removes duplicate elements from the slice.
// It returns a new slice containing only the unique elements from the input slice.
//
// Example:
//   input := []int{1, 2, 2, 3, 4, 4, 5}
//   output := Nub(input) // output will be []int{1, 2, 3, 4, 5}
func Nub[T comparable](slice []T) []T {
    elementMap := make(map[T]bool)
    var newSlice []T
    
    for _, v := range slice {
        if !elementMap[v] {
            elementMap[v] = true
            newSlice = append(newSlice, v)
        }
    }
    return newSlice
}

// Intersection returns a new slice containing elements that are present in both slice1 and slice2.
// It uses a map to track elements in slice1 and checks for their presence in slice2.
//
// Example:
//   slice1 := []int{1, 2, 3, 4}
//   slice2 := []int{3, 4, 5, 6}
//   output := Intersection(slice1, slice2) // output will be []int{3, 4}
func Intersection[T comparable](slice1 []T, slice2 []T) []T {
    elementMap := make(map[T]bool)
    for _, elem := range slice1 {
        elementMap[elem] = true
    }

    var result []T
    for _, elem := range slice2 {
        if elementMap[elem] {
            result = append(result, elem)
        }
    }
    return result
}

// Union returns a new slice containing all unique elements from both slice1 and slice2.
// It uses a map to track unique elements from both slices.
//
// Example:
//   slice1 := []int{1, 2, 3, 4}
//   slice2 := []int{3, 4, 5, 6}
//   output := Union(slice1, slice2) // output will be []int{1, 2, 3, 4, 5, 6}
func Union[T comparable](slice1 []T, slice2 []T) []T {
    elementMap := make(map[T]bool)
    for _, elem := range slice1 {
        elementMap[elem] = true
    }
    for _, elem := range slice2 {
        elementMap[elem] = true
    }

    var result []T
    for elem := range elementMap {
        result = append(result, elem)
    }
    return result
}

// Difference returns a new slice containing elements that are in slice1 but not in slice2.
// It uses a map to track elements in slice2 and checks for their absence in slice1.
//
// Example:
//   slice1 := []int{1, 2, 3, 4}
//   slice2 := []int{3, 4, 5, 6}
//   output := Difference(slice1, slice2) // output will be []int{1, 2}
func Difference[T comparable](slice1 []T, slice2 []T) []T {
    elementMap := make(map[T]bool)
    for _, elem := range slice2 {
        elementMap[elem] = true
    }

    var result []T
    for _, elem := range slice1 {
        if !elementMap[elem] {
            result = append(result, elem)
        }
    }
    return result
}