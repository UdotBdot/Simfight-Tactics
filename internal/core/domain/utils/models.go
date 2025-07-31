package utils

// Property represents a generic characteristic (e.g., trait or role) of a unit.
type Property[T comparable] struct {
	Value T `json:"value"`
}
