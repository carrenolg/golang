package main

func fibCiclo(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibRecursivo(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursivo(n-1) + fibRecursivo(n-2)
}

func main() {}
