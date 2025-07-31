package information

import (
	"errors"
	"fmt"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
)

// NameValidator validates unit names against allowed values
type NameValidator struct {
	validator utils.SetValidator[Name]
}

// NewNameValidator creates a new NameValidator with predefined names
func NewNameValidator() NameValidator {
	return NameValidator{validator: utils.NewSetValidator(Names)}
}

// Validate checks if the name is non-empty and in the allowed set
func (v NameValidator) Validate(name string) error {
	var errs []error
	if name == "" {
		errs = append(errs, errors.New("unit name is required"))
	}
	if !v.validator.IsValid(Name(name)) {
		errs = append(errs, fmt.Errorf("invalid unit name: %s is not in allowed set", name))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// CostValidator validates unit cost values
type CostValidator struct {
	Min int
	Max int
}

// NewCostValidator creates a new CostValidator with given min/max
func NewCostValidator(min, max int) *CostValidator {
	return &CostValidator{Min: min, Max: max}
}

// Validate checks if the cost is within allowed range
func (v *CostValidator) Validate(cost int) error {
	if cost < v.Min || cost > v.Max {
		return fmt.Errorf("invalid unit cost: must be between %d and %d", v.Min, v.Max)
	}
	return nil
}
