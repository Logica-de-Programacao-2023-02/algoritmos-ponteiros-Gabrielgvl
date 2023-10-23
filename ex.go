package main

import (
	"fmt"
	"strings"
)

func main() {
	//contarPalavras("a roupa do rei do japao")
	//countChars("a roupa do rei do japao")

	alunos := map[string][]float64{
		"Joao":    {1, 5, 9},
		"Fulaano": {2.3, 5.6, 8},
		"Tereza":  {9.9, 10, 10},
	}

	medias := calcularMedias(alunos)
	fmt.Println(medias)

}

// a roupa do rei do japao
func contarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
		fmt.Println(ocorrencias)
	}

	return ocorrencias
}

func countChars(text string) map[string]int {
	ocorrencias := make(map[string]int)

	for _, char := range text {
		ocorrencias[string(char)]++
		fmt.Println(ocorrencias)
	}

	return ocorrencias
}

//Escreva uma função que receba um mapa onde as
//chaves são nomes de alunos e os valores são
//slices de notas. A função deve calcular a
//média das notas de cada aluno e retornar um novo
//mapa onde as chaves são os nomes dos alunos e os
//valores são as médias correspondentes.

func calcularMedias(notas map[string][]float64) map[string]float64 {
	resultado := make(map[string]float64)

	for aluno, notasAluno := range notas {
		resultado[aluno] = media(notasAluno)
	}

	return resultado
}

func media(x []float64) float64 {
	soma := 0.
	for _, valor := range x {
		soma += valor
	}
	return soma / float64(len(x))
}
