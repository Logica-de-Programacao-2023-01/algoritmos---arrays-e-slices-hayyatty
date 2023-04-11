package main

import (
	"fmt"
)

func main() {
	var (
		n int
	)
	primes := []int{}
	num := 2
	fmt.Println("Digite um número :")
	fmt.Scan(&n)
	for len(primes) < n {
		isPrime := true

		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
		num++
	}
	fmt.Println(primes)
}
