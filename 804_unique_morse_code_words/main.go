package main

import (
	"strings"
)

var (
	morseITU = map[rune]string{
		rune('a'): ".-",
		rune('b'): "-...",
		rune('c'): "-.-.",
		rune('d'): "-..",
		rune('e'): ".",
		rune('f'): "..-.",
		rune('g'): "--.",
		rune('h'): "....",
		rune('i'): "..",
		rune('j'): ".---",
		rune('k'): "-.-",
		rune('l'): ".-..",
		rune('m'): "--",
		rune('n'): "-.",
		rune('o'): "---",
		rune('p'): ".--.",
		rune('q'): "--.-",
		rune('r'): ".-.",
		rune('s'): "...",
		rune('t'): "-",
		rune('u'): "..-",
		rune('v'): "...-",
		rune('w'): ".--",
		rune('x'): "-..-",
		rune('y'): "-.--",
		rune('z'): "--..",
	}
)

func uniqueMorseRepresentations(words []string) int {
	morseWords := make(map[string]int)

	for i := 0; i < len(words); i++ {
		var sb strings.Builder
		sb.Grow(2)

		for _, v := range words[i] {
			val, found := morseITU[v]
			if found != true {
				panic("not found char in map")
			}
			sb.WriteString(val)
		}

		morseWords[sb.String()] = 0
	}

	return len(morseWords)
}
