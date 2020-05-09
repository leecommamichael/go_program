package level1

import (
	"testing"
)

func TestVariables(t *testing.T) {
	var you int = 0
	var can string = ""
	var name float32 = 12.1
	var these float64 = 12.1
	var almost rune = 'J'
	var anything bool = true
	youCanThinkOf := false

	if !youCanThinkOf {
		t.Error("❌ variables are just the values stored within them")
	}

	if anything {
		t.Error("❌ variables are just the values stored within them")
	}

	if almost == 'J' {
		t.Error("❌ J")
	}

	if these == 12.1 {
		t.Error("❌ these")
	}

	if name == 12.1 {
		t.Error("❌ name")
	}

	if (you + 1) < 5 {
		t.Error("❌ NUM")
	}

	if can+"" == "" {
		t.Error("❌ STR")
	}

}
