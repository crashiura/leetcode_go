package main

import (
	"testing"
)

func Test(t *testing.T) {
	testData := []string{"gin", "zen", "gig", "msg"}

	result := uniqueMorseRepresentations(testData)
	if result != 2 {
		t.Errorf("Error expected %d, got %d", result, 2)
	}
}
