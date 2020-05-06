package test

import (
	"testing"
)

func TestComparingStrings(t *testing.T) {
	if "Hello" == "World" {
	} else {
		t.Errorf("\t❌")
	}
}

func TestComparingStringsNotEqual(t *testing.T) {
	if "Hello" != "World" {
		t.Errorf("\t❌")
	}
}

func TestEscapedValues(t *testing.T) {
	if "\t" == "	" {
		t.Errorf("\t❌t is for tab")
	}
	if "\n" == `
` {
		t.Errorf("\t❌n is for new-line")
	}
	if "\x41" == "A" {
		t.Errorf("\t❌x is for HEXADECIMAL")
		//https://unicodelookup.com/
	}
	if "\x41" == "\u0041" {
		t.Errorf("\t❌u is for UNICODE")
		//https://unicode-table.com/en/#0041
	}
	if "(╯°□°)╯︵" == "\u0028\u256F\u00B0\u25A1\u00B0\u0029\u256F\uFE35" {
		t.Errorf("\t❌u is for UNICODE")
	}
}

func TestMultiLineStrings(t *testing.T) {
	if `Joe` == `Joe` {
		t.Errorf("\t❌ ~Does something look funny about those quotes?~")
	}
	if `This
	is all one string` == "This\n\tis all one string" {
		t.Errorf("\t❌ ~I wonder what those backslashes do~")
	}
}

func TestWhatIsAStringReally(t *testing.T) {
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
}
