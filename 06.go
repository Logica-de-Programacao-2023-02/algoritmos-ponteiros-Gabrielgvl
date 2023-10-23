package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro
//para um struct Livro com campos título e autor, e
//altere o título do livro para "Desconhecido" se o
//autor for "Anônimo".

type Livro struct {
	Titulo string
	Autor  string
}

func main() {
	l1 := Livro{
		Titulo: "ABC",
		Autor:  "FULANO",
	}
	l2 := Livro{
		Titulo: "XXX",
		Autor:  "Anônimo",
	}
	changeBook(&l1)
	changeBook(&l2)

	fmt.Println(l1)
	fmt.Println(l2)
}

func changeBook(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Titulo = "Desconhecido"
	}
}