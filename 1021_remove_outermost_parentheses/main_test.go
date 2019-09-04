package main

import (
	"testing"
)

var testData = map[string]string{
	"(()())(())":         "()()()",
	"(()())(())(()(()))": "()()()()(())",
	"()()":               "",
}

func TestRemoveOuterParentheses(t *testing.T) {
	for k, v := range testData {
		result := removeOuterParentheses(k)
		if result != v {
			t.Errorf("Error expected %s, got %s", result, v)
		}
		t.Log("success")
	}
}
