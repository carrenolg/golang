package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	if resultado := Fib(); resultado != 8 {
		t.Errorf("Fib() = %d ; esperamos 8", resultado)
	}
}

func TestFibCiclo(t *testing.T) {
	pruebas := []struct {
		n        int
		esperado int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, prueba := range pruebas {
		if resultado := FibCiclo(prueba.n); resultado != prueba.esperado {
			t.Errorf("FibCiclo(%d) == %d, esperamos %d", prueba.n, resultado, prueba.esperado)
		}
	}
}

func TestFibRecursivo(t *testing.T) {
	pruebas := []struct {
		n        int
		esperado int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, prueba := range pruebas {
		if resultado := FibRecursivo(prueba.n); resultado != prueba.esperado {
			t.Errorf("FibRecursivo(%d) == %d, esperamos %d", prueba.n, resultado, prueba.esperado)
		}
	}
}

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

// Las funciones del paquete fib pueden interactuar entre sí.
func Example() {
	fmt.Println(FibRecursivo(FibCiclo(5)))
	// Output:
	// 5
}

// FibCiclo recibe un int como argumento, el cual determina la posición en
// la secuencia fibonacci del resultado, también un int. En este ejemplo, para
// la posición 6 en la secuencia fibonacci, obtenemos el número 8.
func ExampleFibCiclo() {
	fmt.Println(FibCiclo(6))
	// Output:
	// 8
}

// FibRecursivo recibe un int como argumento, el cual determina la posición en
// la secuencia fibonacci del resultado, también un int. En este ejemplo, para
// la posición 6 en la secuencia fibonacci, obtenemos el número 8.
func ExampleFibRecursivo() {
	fmt.Println(FibRecursivo(6))
	// Output:
	// 8
}
