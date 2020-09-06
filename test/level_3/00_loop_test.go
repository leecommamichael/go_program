package level3

import "testing"

func TestLoops(t *testing.T) {
	hp := 100

	for hp > 0 {
		hp = hp - 1
		print("\tYou lost 1 HP\n")
	}

	if hp <= 0 {
		t.Errorf("\tYou ran out of HP!\n")
	}
}

func TestSearch(t *testing.T) {
	list := []int{0, 1}

	for i := 0; i < len(list); i++ {
	}
}

func TestSort(t *testing.T) {
	list := []int{0, 1}

	for i := 0; i < len(list); i++ {
	}
}

func TestMap(t *testing.T) {
	list := []int{0, 1}

	for i := 0; i < len(list); i++ {
	}
}

func TestFilter(t *testing.T) {
	list := []int{0, 1}

	for i := 0; i < len(list); i++ {
	}
}

func TestReduce(t *testing.T) {
	list := []int{0, 1}

	for i := 0; i < len(list); i++ {
	}
}
