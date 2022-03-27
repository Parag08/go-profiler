package main

var m = make(map[int]int, 0)

func quickFibonacci(n int) int {
	if n < 2 {
		return n
	}

	var f int
	if v, ok := m[n]; ok {
		f = v
	} else {
		f = fibonacci(n-2) + fibonacci(n-1)
		m[n] = f
	}

	return f
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
