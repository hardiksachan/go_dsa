package utils

func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func SliceEqualWithoutOrder[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		found := false
		for _, w := range b {
			if v == w {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func SliceOfSliceEqual[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !SliceEqual(v, b[i]) {
			return false
		}
	}
	return true
}

func SliceOfSliceEqualWithoutOrder[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		found := false
		for _, w := range b {
			if SliceEqualWithoutOrder(v, w) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
