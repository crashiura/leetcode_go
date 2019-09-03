package main

import (
	"fmt"
	"testing"
)

func Test_one(t *testing.T) {
	t.Log("Test 1 ")
	expected := 10

	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	result := subArraySum(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_two(t *testing.T) {
	t.Log("Test 2 ")
	expected := 2

	nums := []int{1, 2, 3}
	k := 3
	result := subArraySum(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_three(t *testing.T) {
	t.Log("Test 3 ")
	expected := 5

	nums := []int{0, 0, 1, 0, 34, 0, 0, 2}
	k := 0
	result := subArraySum(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_four(t *testing.T) {
	t.Log("Test 4 ")
	expected := 0

	nums := []int{-1, -1, 1}
	k := 0
	result := subArraySum(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_one1(t *testing.T) {
	t.Log("Test 1 ")
	expected := 10

	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	k := 0
	result := subArraySum2(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_two1(t *testing.T) {
	t.Log("Test 2 ")
	expected := 2

	nums := []int{1, 2, 3}
	k := 3
	result := subArraySum2(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}

func Test_there3(t *testing.T) {
	t.Log("Test 2 ")
	expected := 5

	nums := []int{0, 0, 1, 0, 34, 0, 0, 2}
	k := 0
	result := subArraySum2(nums, k)
	fmt.Println(result)

	if expected != result {
		t.Errorf("Expected %v", result)
	}
}
