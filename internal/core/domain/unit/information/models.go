package information

import (
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/traits"
	"github.com/UdotBdot/simfight-tactics/internal/core/domain/unit/role"
	utils "github.com/UdotBdot/simfight-tactics/internal/core/domain/utils"
	"github.com/google/uuid"
)

// Name represents a unit's name
type Name string

// Information contains the core data of a unit
type Information struct {
	Name   string                         `json:"name"`
	Cost   int                            `json:"cost"`
	ID     *uuid.UUID                     `json:"id,omitempty"` // Optional, omitempty for JSON
	Traits []utils.Property[traits.Trait] `json:"traits"`
	Roles  []utils.Property[role.Role]    `json:"roles"`
}
