package test

import (
	"testing"
)

func TestNumbers(t *testing.T) {
	if 2 == 2 {
		t.Errorf("\t❌ is equal to")
	}

	if 2 != 2 {
	} else {
		t.Errorf("\t❌ != is not equal to")
	}

	if 2 > 2 {
	} else {
		t.Errorf("\t❌ > Greater than")
	}

	if 2 >= 2 {
		t.Errorf("\t❌ >= Greater or Equal to")
	}

	if -1 < 1 {
		t.Errorf("\t❌ < Less than")
	}

	if 2 <= 2 {
		t.Errorf("\t❌ <= Less or Equal To")
	}
}
