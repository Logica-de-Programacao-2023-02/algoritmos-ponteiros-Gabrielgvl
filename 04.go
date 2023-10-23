package main

import (
	"fmt"
	"strconv"
)

// Escreva uma função em Go que receba um ponteiro
//para uma variável inteira e atualize o valor
//da variável para a soma dos valores dos seus dois
//últimos dígitos (por exemplo, se o valor inicial da
//variável for 1234, o novo valor será 3+4=7).

func main() {
	x := 1234
	updateDigitsV2(&x)
	fmt.Println(x)
}

func updateDigitsV2(x *int) {
	last := *x % 10
	secondLast := (*x % 100) / 10
	*x = last + secondLast
}

func updateDigits(x *int) {
	xStr := strconv.Itoa(*x)

	last := xStr[len(xStr)-1]
	secondLast := xStr[len(xStr)-2]

	lastI, _ := strconv.Atoi(string(last))
	secondLastI, _ := strconv.Atoi(string(secondLast))

	*x = lastI + secondLastI
}
