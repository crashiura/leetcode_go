package main

import (
	"reflect"
	"testing"
)

var data1 = [][]int{
	[]int{1, 1, 0},
	[]int{1, 0, 1},
	[]int{0, 0, 0},
}

var answer1 = [][]int{
	[]int{1, 0, 0},
	[]int{0, 1, 0},
	[]int{1, 1, 1},
}

var data2 = [][]int{
	[]int{1, 1, 0, 0},
	[]int{1, 0, 0, 1},
	[]int{0, 1, 1, 1},
	[]int{1, 0, 1, 0},
}

var answer2 = [][]int{
	[]int{1, 1, 0, 0},
	[]int{0, 1, 1, 0},
	[]int{0, 0, 0, 1},
	[]int{1, 0, 1, 0},
}

func Test(t *testing.T) {
	result := flipAndInvertImage(data1)
	if compare := compareSlice(result, answer1); !compare {
		t.Errorf("Error expected %v, got %v", result, answer1)
	}

	result = flipAndInvertImage(data2)
	if compare := compareSlice(result, answer2); !compare {
		t.Errorf("Error expected %v, got %v", result, answer2)
	}
}

func compareSlice(a, b [][]int) bool {
	return reflect.DeepEqual(a, b)
}
