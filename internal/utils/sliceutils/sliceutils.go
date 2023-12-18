package sliceutils

import "slices"

// Returns true if the predicate returns true for every element
// of the slice.
func All[T any](input []T, f func(T) bool) bool {
	for _, val := range input {
		if !f(val) {
			return false
		}
	}
	return true
}

// Returns a new slice, keeping all items that the predicate
// returns true for.
func KeepIf[T any](input []T, f func(T) bool) []T {
	output := []T{}
	for _, val := range input {
		if f(val) {
			output = append(output, val)
		}
	}
	return output
}

// Append another slice to the input slice, allocating a new
// slice under the hood if necessary.
func AppendSlice[T any](input []T, more []T) []T {
	totalLen := len(input) + len(more)
	if cap(input) < totalLen {
		input = slices.Grow(input, totalLen)
	}
	for _, item := range more {
		input = append(input, item)
	}
	return input
}

// Map all values in the slice to new values, returning them.
func Map[T any, R any](input []T, f func(T) R) []R {
	outputs := []R{}
	for _, val := range input {
		outputs = append(outputs, f(val))
	}
	return outputs
}

// Fold a slice of values into a single value
func Fold[T any, R any](input []T, r R, f func(R, T) R) R {
	for _, val := range input {
		r = f(r, val)
	}
	return r
}

// Get last element in slice or panic if empty
func Last[T any](input []T) T {
	return input[len(input)-1]
}
