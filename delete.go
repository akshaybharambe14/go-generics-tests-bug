package slices

// Delete removes the element at index i from s. It MUTATES s.
//
// Deleted element is replaced by zero value of T so that it can be garbage collected.
func Delete[T any](s []T, i int) ([]T, bool) {
	len := len(s)
	if i < 0 || i >= len {
		return s, false
	}

	if i < len-1 {
		copy(s[i:], s[i+1:])
	}

	var zero T
	s[len-1] = zero

	return s[:len-1], true
}
