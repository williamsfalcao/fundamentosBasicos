package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrnado na função para verificar se o alunoa esta aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	defer funcao1()
	funcao2()
	fmt.Println(alunoAprovado(7, 8))

}
