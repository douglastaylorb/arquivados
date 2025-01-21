package main

import "fmt"

func main() {

	//for range
	frutas := []string{"banana", "maçã", "uva", "morango"}

	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
