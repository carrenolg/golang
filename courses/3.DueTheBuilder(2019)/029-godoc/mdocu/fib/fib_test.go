package fib

import (
	"fmt"
	"testing"
)

func TestFibCiclo(t *testing.T) {
	n, esperado := 6, 8
	if resultado := FibCiclo(n); resultado != esperado {
		t.Errorf("fibCiclo(%d) == %d, esperamos %d", n, resultado, esperado)
	}
}

func TestFibRecursivo(t *testing.T) {
	n, esperado := 6, 8
	if resultado := FibRecursivo(n); resultado != esperado {
		t.Errorf("fibCiclo(%d) == %d, esperamos %d", n, resultado, esperado)
	}
}

// Benchmark
func BenchmarkFibCiclo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibCiclo(25)
	}
}

func BenchmarkFibRecursivo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursivo(25)
	}
}

// Example
func ExampleFibCiclo() {
	fmt.Println(FibCiclo(6))
	// Output:
	// 8
}

// Las funciones del paquete fib pueden interactuar entre si
func Example() {
	fmt.Println(FibRecursivo(FibCiclo(5)))
	// Output:
	// 5
}

// Fib recurso como example
func ExampleFibRecursivo() {
	fmt.Println(FibRecursivo(5))
	// Output:
	// 5
}
