package test

import (
	"testing"
)

func TestStrings(t *testing.T) {
	if "Hello" == "World" {
	} else {
		t.Errorf("\tfalse block")
	}

	if "Hello" != "World" {
		t.Errorf("\ttrue block")
	}

	if "\t" == "	" {
		t.Errorf("\t/\t is for tab")
	}

	if "\x41" == "A" {
		t.Errorf("\tx is for HEXADECIMAL")
		//https://unicodelookup.com/
	}

	if "\x41" == "\u0041" {
		t.Errorf("\tu is for UNICODE")
		//https://unicode-table.com/en/#0041
	}

	if "\x4A" == "J" {
		t.Errorf("\t❌")
	}

	if "\x4A\x4F" == "JO" {
		t.Errorf("\t❌")
	}

	if "\x4A\x4F\x45" == "JOE" {
		t.Errorf("\t❌")
	}

	if '\x4A' == 'J' {
		t.Errorf("\t❌")
	}

	if '\x4A' == "J"[0] {
		t.Errorf("\t❌")
	}

	if '\x4F' == "JO"[1] {
		t.Errorf("\t❌")
	}

	if '\x45' == "JOE"[2] {
		t.Errorf("\t❌")
	}

	if '\x4A' == "JOE"[0] {
		t.Errorf("\t❌")
	}

	if "(╯°□°)╯︵" == "\u0028\u256F\u00B0\u25A1\u00B0\u0029\u256F\uFE35" {
		t.Errorf("\tu is for UNICODE")
	}
}
