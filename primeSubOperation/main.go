package main

import (
	"fmt"
	"math"
	"slices"
	//"sort"
)

func main() {
	fmt.Println(primeSubOperation([]int{998, 2}))
	fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	fmt.Println(primeSubOperation([]int{5, 8, 3}))
	// fmt.Println(generatePrimes())
}

func primeSubOperation(nums []int) bool {
	maxElem := slices.Max(nums)
	sieve := make([]bool, maxElem+1)
	fill(sieve, true)
	sieve[1] = false
	for i := 2; i <= int(math.Sqrt(float64(maxElem+1))); i++ {
		if sieve[i] {
			for j := i * i; j <= maxElem; j += i {
				sieve[j] = false
			}
		}
	}
	val := 1
	for i := 0; i < len(nums); {
		difference := nums[i] - val
		if difference < 0 {
			return false
		}
		if sieve[difference] || difference == 0 {
			i++
			val++
		} else {
			val++
		}
	}
	return true
}

func fill(arr []bool, val bool) {
	for i := range arr {
		arr[i] = val
	}
}

// SECOND SOLUTION WITH GEN(WORSE)
func primeSubOperationSecond(nums []int) bool {
	primes := generatePrimes()
	p := 1
	for i := 0; i < len(nums); {
		difference := nums[i] - p
		if difference < 0 {
			return false
		}
		isPrime := difference == 0
		if !isPrime {
			_, isPrime = slices.BinarySearch(primes, difference)
		}
		if isPrime {
			i++
			p++
		} else {
			p++
		}
	}
	return true
}

func generatePrimes() []int {
	primes := make([]int, 0, 1000)

	primes = append(primes, 2)
	nextPrime := 3
	for len(primes) < 1000 {
		sqrt := int(math.Sqrt(float64(nextPrime)))
		isPrime := true
		for i := 0; i < len(primes) && primes[i] <= sqrt; i++ {
			if nextPrime%primes[i] == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, nextPrime)
		}
		nextPrime += 2
	}
	return primes
}
