package test

import (
	"testing"
)

func TestTrue(t *testing.T) {
	if true {
		t.Errorf("\t❌ A bool (or boolean) is either true or false")
		//https://en.wikipedia.org/wiki/George_Boole
		//https://en.wikipedia.org/wiki/Boolean_algebra
	}
}

func TestFalse(t *testing.T) {
	if false {
	} else {
		t.Errorf("\t❌ true and false are boolean literals.")
		//https://en.wikipedia.org/wiki/Literal_(computer_programming)
	}
}

func TestNot(t *testing.T) {
	if !false {
		t.Errorf("\t❌ ! means NOT. ! is an operator.")
		//https://en.wikipedia.org/wiki/Operator_(computer_programming)
	}
}

func TestOr(t *testing.T) {
	if true || false {
		t.Errorf("\t❌ || means OR. || is an operator.")
	}
}

func TestAnd(t *testing.T) {
	if true && !false {
		t.Errorf("\t❌ && means AND. && is an operator.")
	}
}

func TestIf(t *testing.T) {
	if true && (false || true) {
		t.Errorf("\t❌ if {} and else {} let you pick block of code to run based on a bool")
		//https://en.wikipedia.org/wiki/Block_(programming)
	}
}
