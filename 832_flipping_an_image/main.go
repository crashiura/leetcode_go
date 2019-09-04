package main

func main() {
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
			// ^ is the binary XOR operator. 0 ^ 1 = 1, 1 ^ 1 = 0
			v[j] = v[j] ^ 1
			v[i] = v[i] ^ 1
		}

		l := len(v)
		if l%2 != 0 {
			v[l/2] = v[l/2] ^ 1
		}
	}

	return A
}
