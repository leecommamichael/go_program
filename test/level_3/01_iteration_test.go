package level3

import "testing"

func TestIteration(t *testing.T) {
	counter := 0
	array := []int{0, 1}

	for i := 0; i < len(array); i++ {
		counter++
	}
	print(7.2 - 1)
	// if counter != TODO {
	// 	t.Errorf("len is a function that tells you how many elements are in an array")
	// }
}
