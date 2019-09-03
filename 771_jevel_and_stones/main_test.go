package main

import (
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	J := "aA"
	S := "aAbbbb"

	result := numJewelsInStones(J, S)
	if result != 3 {
		t.Logf("Error current count %d, got %d", result, 3)
	}
}
