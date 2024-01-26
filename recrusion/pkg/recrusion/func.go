package recrusion

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n < 3 {
		return n
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}
