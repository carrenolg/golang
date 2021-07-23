// fib contiene dos funciones para calcular la secuencia fibinaci
package fib

// FibCiclo calcula la secuencia fibonaci
// implementando un ciclo for
func FibCiclo(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// FibRecursivo calcula la secuencia fibinaci
// implementando una funcion recursiva
func FibRecursivo(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursivo(n-1) + FibRecursivo(n-2)
}
