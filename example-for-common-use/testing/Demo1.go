package test

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-1)
}
