package predicates

import (
	"github.com/luckyluke66/SliceLib/slices"
)

func Elem[T comparable](elem T, slice []T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		} 
	}
	return false
}

func NotElem[T comparable](elem T, slice []T) bool {
	return !Elem(elem, slice)
}

func IsPrefixOf[T comparable](prefix []T, slice []T) bool {
	for i, v := range prefix {
		if v != slice[i] {
			return false
		}
	}
	return true
}

func IsSuffixOf[T comparable](suffix []T, slice []T) bool {
	return IsPrefixOf(slices.Reverse(suffix), slices.Reverse(slice))
}
