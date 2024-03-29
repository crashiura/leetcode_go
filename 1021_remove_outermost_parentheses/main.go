package main

import (
	"strings"
)

func main() {
}

func removeOuterParentheses(S string) string {
	count := 0
	var sb strings.Builder
	sb.Grow(40)

	for _, v := range S {
		if v == rune('(') {
			if count > 0 {
				sb.WriteRune(v)
			}
			count++
		}

		if v == rune(')') {
			count--
			if count > 0 {
				sb.WriteRune(v)
			}
		}
	}

	return sb.String()
}
