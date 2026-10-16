package main

func getLast[T any](s []T) T {
	// ?
	if len(s) == 0 {
		var zero T
		return zero
	}

	length := len(s)
	return s[length-1]
}
