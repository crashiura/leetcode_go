package main

import (
	"strings"
)

func main() {
}

func defangIP(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func defangIPLOOP(address string) string {
	var sb strings.Builder
	sb.Grow(40)

	for _, v := range address {
		if v == '.' {
			sb.WriteString("[.]")
			continue
		}
		sb.WriteRune(v)
	}
	return sb.String()
}
