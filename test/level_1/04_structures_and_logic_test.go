package level1

import (
	"testing"
)

func statIsValid(specialStat int) bool {
	return specialStat > 0 && specialStat < 11
}

func specialIsValid(special Special) bool {
	return statIsValid(special.strength) &&
		statIsValid(special.perception) &&
		statIsValid(special.endurance) &&
		statIsValid(special.charisma) &&
		statIsValid(special.intelligence) &&
		statIsValid(special.agility) &&
		statIsValid(special.luck)
}

func TestFalloutStructuresWithLogic(t *testing.T) {
	var falloutJoe = Special{2, 5, 5, 5, 5, 5, 5}
	if falloutJoe.strength < 3 {
		t.Error("❌ You failed the STR check.")
	}

	if falloutJoe.intelligence < 8 {
		t.Error("❌ You failed the INT check.")
	}

	falloutJoe.intelligence = 0

	if falloutJoe.intelligence == 0 {
		t.Error("❌ You have no INT")
	}

	falloutJoe.luck = 15

	if !specialIsValid(falloutJoe) {
		t.Error("❌ A SPECIAL stat can't be zero, and it can't be larger than 10")
	}
}
