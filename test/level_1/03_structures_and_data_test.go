package level1

import (
	"testing"
)

type Special struct {
	strength     int
	perception   int
	endurance    int
	charisma     int
	intelligence int
	agility      int
	luck         int
}

func TestFalloutStructures(t *testing.T) {
	falloutJoe := Special{2, 5, 5, 5, 5, 5, 5}

	if falloutJoe.strength < 3 {
		t.Error("❌ You failed the strength check.")
	}

	if falloutJoe.intelligence < 8 {
		t.Error("❌ You failed the strength check.")
	}
}
