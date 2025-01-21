package main

import (
	"fmt"
	"strings"
)

//remover vogais de uma string

// func main() {
// 	fmt.Println(Disemvowel("Hello, World!"))
// }

// func Disemvowel(comment string) string {

// 	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

// 	for _, v := range vowels {
// 		comment = strings.ReplaceAll(comment, v, "")
// 	}

// 	return comment
// }

// Deixar primeira letra de cada palavra em maiúscula

func main() {
	fmt.Println(ToJadenCase("teste cada primeira letra em maiusculo"))
}

func ToJadenCase(comment string) string {

	words := strings.Fields(comment)

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, " ")
}
