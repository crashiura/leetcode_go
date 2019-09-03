package main

import "testing"

func Test_one(t *testing.T) {
	currentString := "address = \"1.1.1.1\""
	expected := defangIP(currentString)
	got := "address = \"1[.]1[.]1[.]1\""

	if expected != got {
		t.Errorf("Received %s, got %s", expected, got)
		return
	}

	t.Log("Success")
}

func BenchmarkDefangIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		currentString := "address = \"1.1.1.1\""
		_ = defangIP(currentString)
	}
}

func BenchmarkDefangIPLOOP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		currentString := "address = \"1.1.1.1\""
		_ = defangIPLOOP(currentString)
	}
}
