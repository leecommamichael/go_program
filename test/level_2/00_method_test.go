package level2

import (
	"testing"
)

func TestMethods(t *testing.T) {
	var falloutJoe FalloutCharacter = FalloutCharacter{
		"Joe Biden",
		1,
		Special{5, 5, 5, 5, 5, 5, 5},
		Skills{},
	}

	if falloutJoe.level <= 1 {
		t.Errorf("❌ Joe is only level %d!", falloutJoe.level)
	}

	if falloutJoe.actionPoints() < 10 {
		t.Errorf("❌ Joe only has %d ActionPoints", falloutJoe.actionPoints())
	}

	if falloutJoe.health() < 300 {
		t.Errorf("❌ Joe has %d HP", falloutJoe.health())
	}

	if falloutJoe.carryWeight() < 350 {
		t.Errorf("❌ Joe can't carry 350lbs, he can carry %dlbs", falloutJoe.carryWeight())
	}

	if falloutJoe.unarmedDamage() < 2 {
		t.Errorf("❌ Joe punches for 2 damage, but his unarmed damage is just %4.2f", falloutJoe.unarmedDamage())
	}
}
