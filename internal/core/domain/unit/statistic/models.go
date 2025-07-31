// Package statistic defines models for unit statistics in the game domain.
package statistic

// Statistics represents the statistical attributes of a unit, including values that vary by star levels (e.g., BaseAttackDamage, Health) and mana pool (start, min, max).
type Statistics struct {
	StarLevels           []int     `json:"star_levels"`
	BaseAttackDamage     []float64 `json:"base_attack_damage"`
	BaseAbilityPower     float64   `json:"base_ability_power"`
	AttackSpeed          float64   `json:"attack_speed"`
	Armor                float64   `json:"armor"`
	MagicResist          float64   `json:"magic_resist"`
	Health               []float64 `json:"health"`
	Range                int       `json:"range"`
	AttackDamagePercent  float64   `json:"attack_damage_percent"`
	AbilityPowerPercent  float64   `json:"ability_power_percent"`
	ManaPool             []int     `json:"mana_pool"`
	ManaRegen            float64   `json:"mana_regen"`
	ManaPerAttack        float64   `json:"mana_per_attack"`
	ManaFromDamage       bool      `json:"mana_from_damage"`
	CriticalStrikeChance float64   `json:"critical_strike_chance"`
	CriticalStrikeDamage float64   `json:"critical_strike_damage"`
	Omnivamp             float64   `json:"omnivamp"`
	DamageAmp            float64   `json:"damage_amp"`
	Durability           float64   `json:"durability"`
	TargetPriority       int       `json:"target_priority"`
}
