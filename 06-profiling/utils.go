package main

import "math"

func IsDivisible(n1, n2 int) bool {
	return n1%n2 == 0
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for idx := 2; idx < n; idx++ {
		if IsDivisible(n, idx) {
			return false
		}
	}
	return true
}

func IsPrime2(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for idx := 2; idx <= (n / 2); idx++ {
		if n%idx == 0 {
			return false
		}
	}
	return true
}

func IsPrime3(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	var idx float64
	//nSqrt = math.Sqrt(float64(n))
	for idx = 2; idx <= math.Sqrt(float64(n)); idx++ {
		if n%int(idx) == 0 {
			return false
		}
	}
	return true
}

func IsPrime4(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	var idx float64
	nSqrt := math.Sqrt(float64(n))
	for idx = 2; idx <= nSqrt; idx++ {
		if n%int(idx) == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(start, end int) []int {
	//var primes []int = make([]int, 100)
	var primes [100]int
	var count = 0
	for idx := start; idx <= end; idx++ {
		if IsPrime(idx) {
			primes[count] = idx
			count++
		}
	}
	return primes[:count]
}
