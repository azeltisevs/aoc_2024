package helpers

import "testing"

func TestGCD(t *testing.T) {
	testData := [][]int{
		{94, 22, 2},
		{270, 192, 6},
		{1, 1, 1},
		{7, 1, 1},
		{121, 11, 11},
		{13, 109, 1},
		{200, 56, 8},
		{770, 900, 10},
	}

	for _, data := range testData {
		a, b, r := data[0], data[1], data[2]
		if ans := GCD(a, b); ans != r {
			t.Errorf("GCD of %d and %d is wrong, expected: %d, got: %d", a, b, r, ans)
		}
	}
}
