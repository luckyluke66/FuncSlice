package stack


// Pop removes the first element from the slice and returns the remaining elements.
// It returns a new slice without the first element.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Pop(input) // output will be []int{2, 3, 4}
func Pop[T any](slice []T) []T {
	return slice[1:]
} 

// Push adds a new element to the beginning of the slice and returns the new slice.
// It returns a new slice with the element added at the beginning.
//
// Example:
//   input := []int{2, 3, 4}
//   val := 1
//   output := Push(input, val) // output will be []int{1, 2, 3, 4}
func Push[T any](slice []T, val T) []T {
	newSlice := []T{val}
	return append(newSlice, slice...)
}