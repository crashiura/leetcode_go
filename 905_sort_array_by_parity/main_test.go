package main

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	expected := []int{2, 4, 1, 3}
	result := sortArrayByParity([]int{3, 1, 2, 4})
	if reflect.DeepEqual(expected, result) != true {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
