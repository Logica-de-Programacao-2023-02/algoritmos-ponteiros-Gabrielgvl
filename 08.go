package main

import "fmt"

// Crie uma função que receba um ponteiro para uma
//variável como argumento e modifique o valor da
//variável dentro da função. Certifique-se de que o
//ponteiro aponte para uma área de memória válida
//e que a memória seja liberada após o uso.

func main() {
	x := 20
	err := change(&x)
	fmt.Println(err)
}

func change(x *int) error {
	if x != nil {
		*x = 10
		return nil
	} else {
		return fmt.Errorf("invalid pointer")
	}
}