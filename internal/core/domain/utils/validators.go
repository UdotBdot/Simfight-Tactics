package utils

import (
	"fmt"
)

// SetValidator defines a generic interface for validating values of type T.
// It provides a method to check if a given value is valid within a predefined set.
type SetValidator[T comparable] interface {
	IsValid(value T) bool
}

// Validator defines a generic validation interface
type Validator[T any] interface {
	Validate(value T) error
}

// NewSetValidator creates a generic validator for a set of allowed values.
// If the input list is empty, returns a validator that rejects all values.
func NewSetValidator[T comparable](list []T) SetValidator[T] {
	if len(list) == 0 {
		return &MapBasedValidator[T]{set: make(map[T]struct{})}
	}
	set := make(map[T]struct{}, len(list))
	for _, v := range list {
		set[v] = struct{}{}
	}
	return &MapBasedValidator[T]{set: set}
}

// MapBasedValidator implements SetValidator using a map for O(1) lookups
type MapBasedValidator[T comparable] struct {
	set map[T]struct{}
}

func (s *MapBasedValidator[T]) IsValid(v T) bool {
	_, ok := s.set[v]
	return ok
}

// ValidateAndCollectErrors runs validation and appends any errors to the provided slice
func ValidateAndCollectErrors[T any](errs *[]error, validator Validator[T], value T, wrapMsg string) {
	if err := validator.Validate(value); err != nil {
		if wrapMsg != "" {
			*errs = append(*errs, fmt.Errorf("%s: %w", wrapMsg, err))
		} else {
			*errs = append(*errs, err)
		}
	}
}
