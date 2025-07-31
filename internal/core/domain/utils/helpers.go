// Package utils provides helper functions for common operations in the game domain.
package utils

import (
	"errors"
	"fmt"
)

// LengthConstraints defines min/max length validation rules
type LengthConstraints struct {
	Min int
	Max int
}

// CleanSlice removes zero values from a slice of comparable types, returning the cleaned slice. Returns empty slice if all zero or input nil.
func CleanSlice[T comparable](slice []T) []T {
	if slice == nil {
		return []T{}
	}
	var zero T
	cleaned := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != zero {
			cleaned = append(cleaned, v)
		}
	}
	return cleaned
}

// ValidateSlice checks if all items in a slice are valid by extracting a value and verifying it against a set of allowed values.
func ValidateSlice[T any, U comparable](items []T, validator SetValidator[U], extract func(T) U) error {
	if items == nil {
		return errors.New("input slice cannot be nil")
	}
	for i, item := range items {
		val := extract(item)
		if !validator.IsValid(val) {
			return fmt.Errorf("invalid value at index %d for type %T: %v (not in allowed set)", i, val, val)
		}
	}
	return nil
}

// CheckDuplicates verify value in a slice
func CheckDuplicates[T comparable](slice []T) error {
	if len(slice) == 0 {
		return nil
	}
	seen := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		if _, exists := seen[v]; exists {
			return fmt.Errorf("duplicate value found: %v", v)
		}
		seen[v] = struct{}{}
	}
	return nil
}

// ValidateLength checks if slice length is within constraints
func ValidateLength[T any](slice []T, constraints LengthConstraints, typeName string) error {
	length := len(slice)
	if length < constraints.Min || length > constraints.Max {
		if constraints.Min == constraints.Max {
			return fmt.Errorf("%s must have exactly %d items; got %d", typeName, constraints.Min, length)
		}
		return fmt.Errorf("%s must be between %d and %d items; got %d", typeName, constraints.Min, constraints.Max, length)
	}
	return nil
}

// ConvertSliceToProperties validates and converts a slice into Property slice. Assumes input is cleaned upstream.
func ConvertSliceToProperties[T comparable](cleaned []T, validator SetValidator[T], constraints LengthConstraints, typeName string) ([]Property[T], error) {
	if err := ValidateLength(cleaned, constraints, typeName); err != nil {
		return nil, err
	}

	if err := ValidateSlice(cleaned, validator, func(v T) T { return v }); err != nil {
		return nil, err
	}

	if err := CheckDuplicates(cleaned); err != nil {
		return nil, err
	}

	result := make([]Property[T], len(cleaned))
	for i, v := range cleaned {
		result[i] = Property[T]{Value: v}
	}
	return result, nil
}
