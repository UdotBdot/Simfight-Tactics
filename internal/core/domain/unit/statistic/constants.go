// Package statistic defines default base statistics constants for units in the game domain.
package statistic

// DefaultStats for each units withour Role implemented.
var DefaultStats = struct {
	AbilityPower         float64
	AttackDamagePercent  float64
	AbilityPowerPercent  float64
	ManaRegen            float64
	ManaPerAttack        float64
	ManaFromDamage       bool
	CriticalStrikeChance float64
	CriticalStrikeDamage float64
	Omnivamp             float64
	DamageAmp            float64
	Durability           float64
	TargetPriority       int
}{
	AbilityPower:         0.0,
	AttackDamagePercent:  0.0,
	AbilityPowerPercent:  0.0,
	ManaRegen:            0.0,
	ManaPerAttack:        0.0,
	ManaFromDamage:       false,
	CriticalStrikeChance: 0.25,
	CriticalStrikeDamage: 1.40,
	Omnivamp:             0.0,
	DamageAmp:            0.0,
	Durability:           0.0,
	TargetPriority:       0,
}
