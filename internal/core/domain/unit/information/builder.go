package information

import (
	"os"
	"strconv"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
)

// Builder constructs unit information components
type Builder struct{}

// NewBuilder creates a new Builder instance
func NewBuilder() *Builder {
	return &Builder{}
}

// getIntFromEnvOrDefault retrieves an int from environment variable or falls back to default
func getIntFromEnvOrDefault(envKey string, defaultVal int) int {
	valStr := os.Getenv(envKey)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal // Fallback on error
	}
	return val
}

// CreateValidatorRegistry creates a configured validator registry
func (b *Builder) CreateValidatorRegistry() ValidatorRegistry {
	const (
		MinTraitsPerUnit = 1
		MaxTraitsPerUnit = 3
		RoleLimit        = 1
	)

	nameValidator := NewNameValidator()
	unitMinCost := getIntFromEnvOrDefault("MIN_UNIT_COST", utils.MinUnitCost)
	unitMaxCost := getIntFromEnvOrDefault("MAX_UNIT_COST", utils.MaxUnitCost)
	costValidator := NewCostValidator(unitMinCost, unitMaxCost)
	traitsSetValidator := utils.NewSetValidator(traits.Traits)
	rolesSetValidator := utils.NewSetValidator(role.Roles)
	traitsMin := getIntFromEnvOrDefault("MIN_TRAITS_PER_UNIT", MinTraitsPerUnit)
	traitsMax := getIntFromEnvOrDefault("MAX_TRAITS_PER_UNIT", MaxTraitsPerUnit)
	rolesLimit := getIntFromEnvOrDefault("ROLE_LIMIT", RoleLimit)
	traitsConstraints := utils.LengthConstraints{Min: traitsMin, Max: traitsMax}
	rolesConstraints := utils.LengthConstraints{Min: rolesLimit, Max: rolesLimit}

	return NewEmptyRegistry().
		WithNameValidator(nameValidator).
		WithCostValidator(costValidator).
		WithTraitsSetValidator(traitsSetValidator).
		WithRolesSetValidator(rolesSetValidator).
		WithTraitsConstraints(traitsConstraints).
		WithRolesConstraints(rolesConstraints)
}

// BuildFactory creates a unit information factory with validators
func (b *Builder) BuildFactory() UnitInfoFactory {
	registry := b.CreateValidatorRegistry()
	return New(registry)
}
