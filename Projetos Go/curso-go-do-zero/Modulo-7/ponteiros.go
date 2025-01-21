package main

import "fmt"

func valuePointer(i *int) {
	*i = 10
}

func main() {
	i := 1

	fmt.Println("Valor inicial de i é: ", i)
	fmt.Println("Endereço de memória de i é: ", &i)

	fmt.Println("Chamando a função valuePointer")
	valuePointer(&i)
	fmt.Println("Valor de i após a função valuePointer é: ", i)
	fmt.Println("Endereço de memória de i após a função valuePointer é: ", &i)
}
