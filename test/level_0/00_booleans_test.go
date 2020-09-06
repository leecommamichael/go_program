package test

import (
	"testing"
)

func TestBooleans(t *testing.T) {
	if true {
		t.Error("\t\tA bool (or boolean) is either true or false")
	} else {
		//this block of code is empty
	}

	if false {
		//this block of code is empty
	} else {
		t.Error("\ttrue and false are boolean values.")
	}

	if !false {
		t.Error("\t! means NOT. ! is an operator.")
	}

	if true || false {
		t.Error("\t|| means OR. || is an operator.")
	}

	if true && !false {
		t.Error("\t&& means AND. && is an operator.")
	}

	if true && (false || true) {
		t.Error("\tif {} and else {} let you pick block of code to run based on a bool")
	}
}
