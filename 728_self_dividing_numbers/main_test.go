package main

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
	result := selfDividingNumbers(1, 22)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error expected %v, got %v", expected, result)
	}
}
