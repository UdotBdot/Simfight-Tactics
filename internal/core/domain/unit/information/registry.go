package information

import (
	"errors"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
)

// NameValidatorRegistry provides name validation
type NameValidatorRegistry interface {
	ValidateName(name string) error
}

// CostValidatorRegistry provides cost validation
type CostValidatorRegistry interface {
	ValidateCost(cost int) error
}

// TraitsValidatorRegistry provides traits set validation and constraints
type TraitsValidatorRegistry interface {
	TraitsSetValidator() utils.SetValidator[traits.Trait]
	TraitsConstraints() utils.LengthConstraints
}

// RolesValidatorRegistry provides roles set validation and constraints
type RolesValidatorRegistry interface {
	RolesSetValidator() utils.SetValidator[role.Role]
	RolesConstraints() utils.LengthConstraints
}

// ValidatorRegistry combines all sub-registries
type ValidatorRegistry interface {
	NameValidatorRegistry
	CostValidatorRegistry
	TraitsValidatorRegistry
	RolesValidatorRegistry
}

// Registry implements ValidatorRegistry with concrete validators
type Registry struct {
	nameValidator      utils.Validator[string]
	costValidator      utils.Validator[int]
	traitsSetValidator utils.SetValidator[traits.Trait]
	rolesSetValidator  utils.SetValidator[role.Role]
	traitsConstraints  utils.LengthConstraints
	rolesConstraints   utils.LengthConstraints
}

// NewEmptyRegistry creates an empty registry instance
func NewEmptyRegistry() *Registry {
	return &Registry{}
}

// WithNameValidator sets the name validator and returns the registry for chaining
func (r *Registry) WithNameValidator(v utils.Validator[string]) *Registry {
	r.nameValidator = v
	return r
}

// WithCostValidator sets the cost validator and returns the registry for chaining
func (r *Registry) WithCostValidator(v utils.Validator[int]) *Registry {
	r.costValidator = v
	return r
}

// WithTraitsSetValidator sets the traits set validator and returns the registry for chaining
func (r *Registry) WithTraitsSetValidator(v utils.SetValidator[traits.Trait]) *Registry {
	r.traitsSetValidator = v
	return r
}

// WithRolesSetValidator sets the roles set validator and returns the registry for chaining
func (r *Registry) WithRolesSetValidator(v utils.SetValidator[role.Role]) *Registry {
	r.rolesSetValidator = v
	return r
}

// WithTraitsConstraints sets the traits constraints and returns the registry for chaining
func (r *Registry) WithTraitsConstraints(c utils.LengthConstraints) *Registry {
	r.traitsConstraints = c
	return r
}

// WithRolesConstraints sets the roles constraints and returns the registry for chaining
func (r *Registry) WithRolesConstraints(c utils.LengthConstraints) *Registry {
	r.rolesConstraints = c
	return r
}

// ValidateName validates the unit name
func (r *Registry) ValidateName(name string) error {
	if r.nameValidator == nil {
		return errors.New("name validator not configured")
	}
	return r.nameValidator.Validate(name)
}

// ValidateCost validates the unit cost
func (r *Registry) ValidateCost(cost int) error {
	if r.costValidator == nil {
		return errors.New("cost validator not configured")
	}
	return r.costValidator.Validate(cost)
}

// TraitsSetValidator returns the traits set validator
func (r *Registry) TraitsSetValidator() utils.SetValidator[traits.Trait] {
	return r.traitsSetValidator
}

// RolesSetValidator returns the roles set validator
func (r *Registry) RolesSetValidator() utils.SetValidator[role.Role] {
	return r.rolesSetValidator
}

// TraitsConstraints returns the traits length constraints
func (r *Registry) TraitsConstraints() utils.LengthConstraints {
	return r.traitsConstraints
}

// RolesConstraints returns the roles length constraints
func (r *Registry) RolesConstraints() utils.LengthConstraints {
	return r.rolesConstraints
}
