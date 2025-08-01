package main

import (
	"fmt"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/information"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
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

	unit, err := factory.Create(
		unitName,
		unitCost,
		unitTraits,
		unitRole,
	)

	if utils.HandleUnitCreationError(err) {
		return
	}
	fmt.Printf("Unit created: %+v\n", unit)
}
