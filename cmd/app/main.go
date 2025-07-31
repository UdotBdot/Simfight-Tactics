package main

import (
	"fmt"
	"log"

	"github.com/UdotBdot/simfight-tactics/internal/core/domain/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/information"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present (ignore error if not found)
	_ = godotenv.Load()

	builder := information.NewBuilder()
	factory := builder.BuildFactory()

	unit, err := factory.Create(
		"Naafiri",
		4,
		[]traits.Trait{"Battle Academia", "Crystal Gambit"},
		[]role.Role{"Attack Tank"},
	)

	if err != nil {
		log.Printf("Unit creation failed: %v", err)
		fmt.Printf("Error during unit initialization: %v\n", err)
		return
	}
	fmt.Printf("Unit created: %+v\n", unit)
}
