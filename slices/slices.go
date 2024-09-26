package slices

// Number is a type constraint that allows float64, int, and float32 types.
type Number interface {
	float64 | int | float32
}

// Pair represents a pair of values of any types.
type Pair[T any, S any] struct {
	Fst T // Fst is the first value in the pair.
    Snd S // Snd is the second value in the pair.
}

// Map applies a function to each element of the slice and returns a new slice with the results.
//
// Example:
//   input := []int{1, 2, 3}
//   output := Map(input, func(x int) int { return x * 2 }) // output will be []int{2, 4, 6}
func Map[T any](slice []T, fn func(T) T) []T {
	mapped := make([]T, len(slice))
	for i, v := range slice {
		mapped[i] = fn(v)
	}
	return mapped
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Filter(input, func(x int) bool { return x%2 == 0 }) // output will be []int{2, 4}
func Filter[T any](slice []T, fn func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Foldr reduces the slice from the right using the provided function and initial value.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Foldr(input, 0, func(x, y int) int { return x + y }) // output will be 10
func Foldr[T any](slice []T, unit T, fn func(T, T) T) T {
	reduced := unit
	for _, v := range slice {
		reduced = fn(reduced, v)
	}

	return reduced
}

// Unfoldr generates a slice by repeatedly applying a function to a starting value.
//
// Example:
//   output := Unfoldr(1, func(x int) int { return x + 1 }, 5) // output will be []int{1, 2, 3, 4, 5}
func Unfoldr[T any](start T, fn func(T) T, len int) []T {
	slice := []T{start}
	n := start
	for i := 0; i < len; i++ {
		n = fn(n)
		slice = append(slice, n)
	}
	return slice
} 

// Intersperse inserts the given element between each element of the input slice.
// It returns a new slice with the interspersed elements.
//
// Example:
//   input := []int{1, 2, 3}
//   elem := 0
//   output := Intersperse(input, elem) // output will be []int{1, 0, 2, 0, 3}
func Intersperse[T any](slice []T, elem T) []T {
	len := len(slice)
	if len == 0 {
        return slice
    }
	var newSlice []T
	for i, v := range slice {
		newSlice = append(newSlice, v)
		if i < len - 1 {
			newSlice = append(newSlice, elem)
		}
	}

	return newSlice
}

func zip[T any, S any](slice1 []T, slice2 []S, len int) []Pair[T, S] {
	var newSlice []Pair[T, S]

	for i := 0; i < len; i++ {
		var pair = Pair[T, S]{slice1[i], slice2[i]}
		newSlice = append(newSlice, pair)
	}

	return newSlice
}

// Zip zips two slices into a slice of pairs. It zips up to the length of the shorter slice.
// It returns a new slice of pairs.
//
// Example:
//   slice1 := []int{1, 2, 3}
//   slice2 := []string{"a", "b", "c", "d"}
//   output := Zip(slice1, slice2) // output will be []Pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}
func Zip[T any, S any](slice1 []T, slice2 []S) []Pair[T, S] {
	len := min(len(slice1), len(slice2))
	return zip(slice1, slice2, len)
}

// Concat concatenates a slice of slices into a single slice.
// It returns a new slice containing all the elements of the input slices.
//
// Example:
//   input := [][]int{{1, 2}, {3, 4}, {5}}
//   output := Concat(input) // output will be []int{1, 2, 3, 4, 5}
func Concat[T any](slices [][]T) []T {
	var newSlice []T
	return Foldr(slices, newSlice, func(s1 []T, s2 []T) []T {return append(s1, s2...)})
}

// And returns true if all elements in the slice are true.
// It returns false if any element is false.
//
// Example:
//   input := []bool{true, true, true}
//   output := And(input) // output will be true
func And(slice []bool) bool {
	return Foldr(slice, true, func(x bool, y bool) bool { return x && y })
}

// Or returns true if any element in the slice is true.
// It returns false if all elements are false.
//
// Example:
//   input := []bool{false, true, false}
//   output := Or(input) // output will be true
func Or(slice []bool) bool {
	return Foldr(slice, false, func(x bool, y bool) bool { return x || y })
}

// Sum returns the sum of all elements in the slice.
// It uses the Foldr function to accumulate the sum.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Sum(input) // output will be 10
func Sum[T Number](slice []T) T {
	return Foldr(slice, 0, func(x T, y T) T { return x + y })
}

// Product returns the product of all elements in the slice.
// It uses the Foldr function to accumulate the product.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Product(input) // output will be 24
func Product[T Number](slice []T) T {
	return Foldr(slice, 1, func(x T, y T) T { return x * y })
} 

// Reverse returns a new slice with the elements of the input slice in reverse order.
//
// Example:
//   input := []int{1, 2, 3, 4}
//   output := Reverse(input) // output will be []int{4, 3, 2, 1}
func Reverse[T any](slice []T) []T {
	var newSlice []T
	len := len(slice) - 1
	for i := len; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}