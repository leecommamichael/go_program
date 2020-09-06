package test

import (
	"testing"
)

func helloWorld() string {
	return "🎉 Hello World! 🎉\n"
}

func ifYouDontHaveSomethingNiceToSay() string {
	return ""
	return "😒 I hate people 😒\n"
}

func two() int {
	return 2
}

func pi() float32 {
	return 3.14
}

func oneTwoThree() []int {
	return []int{1, 2, 3}
}

func giveMeBackMy(someInteger int) int {
	return someInteger
}

func addOneTo(anyNumber int) int {
	return anyNumber + 1
}

func add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

func alwaysTrue() bool {
	return true
}

func onlyTrueIfItsOverTwo(it int) bool {
	return it > 2
}

func TestFunctions(t *testing.T) {
	if helloWorld() == "🎉 Hello World! 🎉\n" {
		t.Errorf("\t❌ writing () after a function's name will \"call\" that function. It will return to you a value.")
	}

	if ifYouDontHaveSomethingNiceToSay() == "" {
		t.Errorf("\t❌ a function stops running code when it returns a value to the caller.")
	}

	if two() == 2 {
		t.Errorf("\t❌ a function has a return-type. It will always returns that kind of value to the caller. This one's type is an int.")
	}

	if pi() != 3.15 {
		t.Errorf("\t❌")
	}

	if oneTwoThree()[0] == 1 {
		t.Errorf("\t❌")
	}

	if alwaysTrue() {
		t.Errorf("\t❌")
	}

	if giveMeBackMy(2) == 2 {
		t.Errorf("\t❌")
	}

	if addOneTo(2) != 4 {
		t.Errorf("\t❌")
	}

	if add(2, 2) != 5 {
		t.Errorf("\t❌")
	}

	if onlyTrueIfItsOverTwo(3) {
		t.Errorf("\t❌")
	}
}
