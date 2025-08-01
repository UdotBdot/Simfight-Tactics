// Package role defines constants for unit role behaviors in the game domain.
// Specialist Role not handled due to different behavior.
package role

// backup for .env value
const (
	RoleLimit = 1
)

// Constants defines the behavioral parameters for a unit role
type Constants struct {
	ManaPerAttack  float64
	ManaFromDamage bool
	ManaRegen      float64
	TargetPriority int
	Omnivamp       float64
}

// RoleConstants maps each role to its behavioral constants
var RoleConstants = map[Role]Constants{
	"Attack Tank": {
		ManaPerAttack:  5.0,
		ManaFromDamage: true,
		TargetPriority: 1,
	},
	"Magic Tank": {
		ManaPerAttack:  5.0,
		ManaFromDamage: true,
		TargetPriority: 1,
	},
	"Hybrid Tank": {
		ManaPerAttack:  5.0,
		ManaFromDamage: true,
		TargetPriority: 1,
	},
	"Attack Fighter": {
		ManaPerAttack: 10.0,
		Omnivamp:      0.1,
	},
	"Magic Fighter": {
		ManaPerAttack: 10.0,
		Omnivamp:      0.1,
	},
	"Hybrid Fighter": {
		ManaPerAttack: 10.0,
		Omnivamp:      0.1,
	},
	"Attack Assassin": {
		ManaPerAttack:  10.0,
		TargetPriority: -1,
	},
	"Magic Assassin": {
		ManaPerAttack:  10.0,
		TargetPriority: -1,
	},
	"Hybrid Assassin": {
		ManaPerAttack:  10.0,
		TargetPriority: -1,
	},
	"Attack Marksman": {
		ManaPerAttack: 10.0,
	},
	"Magic Marksman": {
		ManaPerAttack: 10.0,
	},
	"Hybrid Marksman": {
		ManaPerAttack: 10.0,
	},
	"Attack Caster": {
		ManaPerAttack: 7.0,
		ManaRegen:     2.0,
	},
	"Magic Caster": {
		ManaPerAttack: 7.0,
		ManaRegen:     2.0,
	},
	"Hybrid Caster": {
		ManaPerAttack: 7.0,
		ManaRegen:     2.0,
	},
	"Attack Specialist": {},
	"Magic Specialist":  {},
	"Hybrid Specialist": {},
}
