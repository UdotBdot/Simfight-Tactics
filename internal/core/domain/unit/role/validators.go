package role

import (
	"fmt"
)

// RoleLimit defines the exact number of roles required per unit
const (
	RoleLimit = 1
)

// Validator validates role collections
type Validator struct{}

// Validate checks if the role collection contains exactly one role
func (v Validator) Validate(roles []Role) error {
	length := len(roles)
	if length != RoleLimit {
		return fmt.Errorf("unit must have exactly one role; got %d", length)
	}
	return nil
}
