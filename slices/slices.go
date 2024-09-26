package slices

type Number interface {
	float64 | int | float32
}

type Pair[T any, S any] struct {
	Fst T
	Snd S
}


func Map[T any](slice []T, fn func(T) T) []T {
	mapped := make([]T, len(slice))
	for i, v := range slice {
		mapped[i] = fn(v)
	}
	return mapped
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var filtered []T
	for _, v := range slice {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Foldr[T any](slice []T, unit T, fn func(T, T) T) T {
	reduced := unit
	for _, v := range slice {
		reduced = fn(reduced, v)
	}

	return reduced
}

func Unfoldr[T any](start T, fn func(T) T, len int) []T {
	slice := []T{start}
	n := start
	for i := 0; i < len; i++ {
		n = fn(n)
		slice = append(slice, n)
	}
	return slice
} 

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

func Zip[T any, S any](slice1 []T, slice2 []S) []Pair[T, S] {
	len := min(len(slice1), len(slice2))
	return zip(slice1, slice2, len)
}

func Concat[T any](slices [][]T) []T {
	var newSlice []T
	return Foldr(slices, newSlice, func(s1 []T, s2 []T) []T {return append(s1, s2...)})
}

func And(slice []bool) bool {
	return Foldr(slice, true, func(x bool, y bool) bool { return x && y })
}

func Or(slice []bool) bool {
	return Foldr(slice, false, func(x bool, y bool) bool { return x || y })
}

func Any[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	
	return false
}

func All[T any](slice []T, fn func(T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Sum[T Number](slice []T) T {
	return Foldr(slice, 0, func(x T, y T) T { return x + y })
}

func Product[T Number](slice []T) T {
	return Foldr(slice, 1, func(x T, y T) T { return x * y })
} 

func Reverse[T any](slice []T) []T {
	var newSlice []T
	len := len(slice) - 1
	for i := len; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}