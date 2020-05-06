package test

import (
	"testing"
)

func TestComparingNumbersEqual(t *testing.T) {
	if 2 == 2 {
		t.Errorf("\t❌ is equal to")
	}
}

func TestComparingNumbersNotEqual(t *testing.T) {
	if 2 != 2 {
	} else {
		t.Errorf("\t❌ != is not equal to")
	}
}

func TestGreater(t *testing.T) {
	if 2 > 2 {
	} else {
		t.Errorf("\t❌ > Greater than")
	}
}

func TestGreaterOrEqual(t *testing.T) {
	if 2 >= 2 {
		t.Errorf("\t❌ >= Greater or Equal to")
	}
}

func TestLess(t *testing.T) {
	if -1 < 1 {
		t.Errorf("\t❌ < Less than")
	}
}

func TestLessOrEqual(t *testing.T) {
	if 2 <= 2 {
		t.Errorf("\t❌ <= Less or Equal To")
	}
}
