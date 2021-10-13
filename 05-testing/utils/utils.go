package utils

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for idx := 2; idx < n; idx++ {
		if n%idx == 0 {
			return false
		}
	}
	return true
}
