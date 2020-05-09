package level1

import (
	"testing"
)

type Person struct {
	firstName string
	lastName  string
}

func TestStructures(t *testing.T) {
	joe := Person{"Joe", "Schleipman"}
	someCrap := Person{"01234", "whatever"}
	var emptyPerson = Person{}

	if joe.firstName == "Joe" {
		t.Error("❌ ")
	}
	if (Person{"Joe", "Schlippy"}).firstName == "Joe" {
		t.Error("❌ ")
	}
	if emptyPerson.firstName == "" {
		t.Error("❌ ")
	}
	if emptyPerson.lastName == "" {
		t.Error("❌ ")
	}
	if someCrap.firstName == "whatever" {
		t.Error("❌ ")
	}
}
