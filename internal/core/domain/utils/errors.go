package utils

import (
	"fmt"
)

// HandleUnitCreationError prints the error if present and returns true if an error occurred (for early return).
func HandleUnitCreationError(err error) bool {
	if err != nil {
		fmt.Printf("Error during unit initialization: %v\n", err)
		return true
	}
	return false
}
