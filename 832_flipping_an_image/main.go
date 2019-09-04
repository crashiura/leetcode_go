package main

func main() {
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]

			if v[j] == 0 {
				v[j] = 1
			} else {
				v[j] = 0
			}
			if v[i] == 0 {
				v[i] = 1
			} else {
				v[i] = 0
			}
		}

		l := len(v)
		if l%2 != 0 {
			if v[l/2] == 0 {
				v[l/2] = 1
			} else {
				v[l/2] = 0
			}
		}
	}

	return A
}
