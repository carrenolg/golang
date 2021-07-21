package main

import (
	"testing"
)

func TestFibCiclo(t *testing.T) {
	n, esperado := 6, 8
	if resultado := fibCiclo(n); resultado != esperado {
		t.Errorf("fibCiclo(%d) == %d, esperamos %d", n, resultado, esperado)
	}
}

func TestFibRecursivo(t *testing.T) {
	n, esperado := 6, 8
	if resultado := fibRecursivo(n); resultado != esperado {
		t.Errorf("fibCiclo(%d) == %d, esperamos %d", n, resultado, esperado)
	}
}

// Benchmark
func BenchmarkFibCiclo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibCiclo(25)
	}
}

func BenchmarkFibRecursivo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibRecursivo(25)
	}
}
