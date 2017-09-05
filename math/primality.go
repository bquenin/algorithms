package math

import "math"

var primes = map[int64]bool{
	2:  true,
	3:  true,
	5:  true,
	7:  true,
	11: true,
}

func IsPrime(n int64) bool {
	if n < 2 {
		return false
	}
	_, isPrime := primes[n]
	if isPrime {
		return true
	}
	for prime := range primes {
		if n%prime == 0 {
			return false
		}
	}
	for i := int64(11); i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	primes[n] = true
	return true
}
