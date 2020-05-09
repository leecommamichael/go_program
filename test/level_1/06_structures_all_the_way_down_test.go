package level1

import (
	"math"
	"testing"
)

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

type DerivedStats struct {
	health         int
	actionPoints   int
	carryWeight    int
	meleeDamage    float32
	unarmedDamage  float32
	criticalChance float32
	// damageResistance    int
	// fireResistance      int
	// poisonResistance    int
	// radiationResistance int
}

func deriveStats(character FalloutCharacter) DerivedStats {
	return DerivedStats{
		90 + (character.stats.endurance * 20) + (character.level * 10),
		5 + int(math.Floor(float64(character.stats.agility/2.0))),
		150 + character.stats.strength*10,
		float32(character.stats.strength) * 1.5,
		(float32(character.skills.unarmed) / 20) + 0.5,
		float32(character.stats.luck) * 0.01,
	}
}

type FalloutCharacter struct {
	name   string
	level  int
	stats  Special
	skills Skills
}

func TestFalloutCharacters(t *testing.T) {
	var falloutJoe FalloutCharacter = FalloutCharacter{
		"Joe Biden",
		1,
		Special{5, 5, 5, 5, 5, 5, 5},
		Skills{},
	}

	if falloutJoe.level <= 1 {
		t.Errorf("❌ Joe is only level %d!", falloutJoe.level)
	}

	var moreJoeBidenStats DerivedStats = deriveStats(falloutJoe)

	if moreJoeBidenStats.actionPoints < 10 {
		t.Errorf("❌ Joe only has %d ActionPoints", moreJoeBidenStats.actionPoints)
	}

	if moreJoeBidenStats.health < 300 {
		t.Errorf("❌ Joe has %d HP", moreJoeBidenStats.health)
	}

	if moreJoeBidenStats.carryWeight < 350 {
		t.Errorf("❌ Joe can't carry 350lbs, he can carry %dlbs", moreJoeBidenStats.carryWeight)
	}

	if moreJoeBidenStats.unarmedDamage < 2 {
		t.Errorf("❌ Joe punches for 2 damage, but his unarmed damage is just %4.2f", moreJoeBidenStats.unarmedDamage)
	}
}
