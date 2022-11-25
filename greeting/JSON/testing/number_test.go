package testing

import "testing"

func TestSumOfNumbers(t *testing.T) {
	col := SumOfNumbers(1, 2, 3, 4, 5)
	result := 15
	if col != result {
		t.Error(``)
	}

}
