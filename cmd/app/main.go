package main

import (
	"fmt"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/information"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/traits"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present (ignore error if not found)
	_ = godotenv.Load()

	builder := information.NewBuilder()
	factory := builder.BuildFactory()
	unitName := "Naafiri"
	unitCost := 4
	unitTraits := []traits.Trait{"Soul Fighter", "Juggernaut"}
	unitRole := []role.Role{"Attack Tank"}

	// quick object init, TODO: test
	unit, err := factory.Create(
		unitName,
		unitCost,
		unitTraits,
		unitRole,
	)

	if handleUnitCreationError(err) {
		return
	}
	fmt.Printf("Unit created: %+v\n", unit)
}

// handleUnitCreationError prints the error if present and returns true if an error occurred (for early return).
func handleUnitCreationError(err error) bool {
	if err != nil {
		fmt.Printf("Error during unit initialization: %v\n", err)
		return true
	}
	return false
}
