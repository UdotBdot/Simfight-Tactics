package information

import (
	"errors"
	"fmt"
	"log"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	utils "github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
	"github.com/google/uuid"
)

// UnitInfoFactory defines the contract for creating unit information
type UnitInfoFactory interface {
	Create(name string, cost int, traits []traits.Trait, roles []role.Role) (*Information, error)
}

// Factory implements UnitInfoFactory with validation
type Factory struct {
	registry ValidatorRegistry
}

// New creates a new Factory with the provided validator registry
func New(registry ValidatorRegistry) UnitInfoFactory {
	return &Factory{registry: registry}
}

// Create builds a new Information instance after validating all inputs
func (f *Factory) Create(name string, cost int, traitsParam []traits.Trait, rolesParam []role.Role) (*Information, error) {
	var errs []error

	if err := f.registry.ValidateName(name); err != nil {
		log.Printf("Validation error for name '%s': %v", name, err)
		errs = append(errs, fmt.Errorf("invalid name: %w", err))
	}

	if err := f.registry.ValidateCost(cost); err != nil {
		log.Printf("Validation error for cost %d: %v", cost, err)
		errs = append(errs, fmt.Errorf("invalid cost: %w", err))
	}

	cleanedTraits := utils.CleanSlice(traitsParam)
	validatedTraits, err := utils.ConvertSliceToProperties(cleanedTraits, f.registry.TraitsSetValidator(), f.registry.TraitsConstraints(), "traits")
	if err != nil {
		log.Printf("Validation error for traits %v: %v", traitsParam, err)
		errs = append(errs, fmt.Errorf("invalid traits: %w", err))
	}

	cleanedRoles := utils.CleanSlice(rolesParam)
	validatedRoles, err := utils.ConvertSliceToProperties(cleanedRoles, f.registry.RolesSetValidator(), f.registry.RolesConstraints(), "roles")
	if err != nil {
		log.Printf("Validation error for roles %v: %v", rolesParam, err)
		errs = append(errs, fmt.Errorf("invalid roles: %w", err))
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	var id *uuid.UUID // Optional; set to nil by default
	// If persistence needed, generate: id = uuidPtr(uuid.New())

	return &Information{
		Name:   name,
		Cost:   cost,
		ID:     id,
		Traits: validatedTraits,
		Roles:  validatedRoles,
	}, nil
}

// e.g:
// id := uuid.New()
// info.ID = uuidPtr(id)
// // Convert uuid.UUID in *uuid.UUID

// uuidPtr helper for pointer
// func uuidPtr(u uuid.UUID) *uuid.UUID {
// 	return &u
// }
