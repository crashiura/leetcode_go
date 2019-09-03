package main

import (
	"fmt"
	"strings"
)

func main() {
	res := numJewelsInStones("aA", "aAAbbbb")
	fmt.Println(res)
}

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, v := range S {
		if strings.ContainsAny(J, string(v)) {
			count++
		}
	}

	return count
}
