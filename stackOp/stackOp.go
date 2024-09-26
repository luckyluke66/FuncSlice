package stackOp

func Pop[T any](slice []T) []T {
	return slice[1:]
} 

func Push[T any](slice []T, val T) []T {
	newSlice := []T{val}
	return append(newSlice, slice...)
}