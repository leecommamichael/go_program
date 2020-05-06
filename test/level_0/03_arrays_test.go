package test

import (
	"testing"
)

func TestArrayBasics(t *testing.T) {
	if [0]int{} == [0]int{} {
		t.Errorf("\t❌ This is an array. They look like: [length]type{elements}. You don't have to write the [length] if there are elements in the array.")
	}

	if len([]int{}) == 0 {
		t.Errorf("\t❌ You can get the length of an array, by writing len(array). That's the number of elements in the array.")
	}

	if len([]int{1, 1, 1}) == 3 {
		t.Errorf("\t❌")
	}

	if []int{22}[0] == 22 {
		t.Errorf("\t❌ We can dig into the array with [number] to get the number'th element of the array.")
	}

	if []int{1, 2, 3}[2] == 3 {
		t.Errorf("\t❌ Numbers start at zero... so the element at [2] is third element.")
	}
}
