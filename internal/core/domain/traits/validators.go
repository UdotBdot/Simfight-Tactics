package traits

import (
	"fmt"
)

// This const defines the range of available Traits for each unit.
const (
	MinTraitsPerUnit = 1
	MaxTraitsPerUnit = 3
)

// Validator validates trait collections
type Validator struct{}

// Validate checks if the trait collection meets the required constraints
func (v Validator) Validate(traits []Trait) error {
	length := len(traits)
	if length < MinTraitsPerUnit || length > MaxTraitsPerUnit {
		return fmt.Errorf("unit traits must be between %d and %d; got %d", MinTraitsPerUnit, MaxTraitsPerUnit, length)
	}
	return nil
}
