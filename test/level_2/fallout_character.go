package level2

import "math"

// Special ...
type Special struct {
	strength     int
	perception   int
	endurance    int
	charisma     int
	intelligence int
	agility      int
	luck         int
}

// Skills ...
type Skills struct {
	bigGuns       int
	energyWeapons int
	explosives    int
	meleeWeapons  int
	smallGuns     int
	unarmed       int
	lockpick      int
	medicine      int
	repair        int
	science       int
	sneak         int
	barter        int
	speech        int
}

// FalloutCharacter ...
type FalloutCharacter struct {
	name   string
	level  int
	stats  Special
	skills Skills
}

func (character FalloutCharacter) health() int {
	return 90 + (character.stats.endurance * 20) + (character.level * 10)
}
func (character FalloutCharacter) actionPoints() int {
	return 5 + int(math.Floor(float64(character.stats.agility/2.0)))
}
func (character FalloutCharacter) carryWeight() int {
	return 150 + character.stats.strength*10
}
func (character FalloutCharacter) meleeDamage() float32 {
	return float32(character.stats.strength) * 1.5
}
func (character FalloutCharacter) unarmedDamage() float32 {
	return (float32(character.skills.unarmed) / 20) + 0.5
}
func (character FalloutCharacter) criticalChance() float32 {
	return float32(character.stats.luck) * 0.01
}
