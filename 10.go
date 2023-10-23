package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma slice e seu tamanho e
//preencha-o com os n primeiros números primos.

func main() {
	var s []int
	fillPrimes(&s, 10)
	fmt.Println(s)
}

func fillPrimes(s *[]int, n int) {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			*s = append(*s, i)
		}
	}
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
