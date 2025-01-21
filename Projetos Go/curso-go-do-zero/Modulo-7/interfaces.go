package main

// import "fmt"

// type Cachorro struct {
// 	Raca  string
// 	Cor   string
// 	Patas int
// }

// type Gato struct {
// 	Cor   string
// 	Patas int
// }

// type Animal interface {
// 	Barulho() string
// 	NumeroDePatas() int
// }

// func (c Cachorro) Barulho() string {
// 	return "Au au"
// }

// func (g Gato) Barulho() string {
// 	return "Miau"
// }

// func (c Cachorro) NumeroDePatas() int {
// 	return c.Patas
// }

// func (g Gato) NumeroDePatas() int {
// 	return g.Patas
// }

// func FazBarulho(a Animal) {
// 	fmt.Println(a.Barulho())
// }

// type Pessoa struct {
// 	Nome  string
// 	Idade int
// }

// type MaiorDeIdade struct {
// 	Pessoa
// 	Profissao string
// }

// type MenorDeIdade struct {
// 	Pessoa
// 	AnoColegial int
// }

// type Adulto interface {
// 	Trabalhar() string
// }

// type Crianca interface {
// 	Estudar() string
// }

// func (ma MaiorDeIdade) Trabalhar() string {
// 	return fmt.Sprintf("%s trabalha como %s", ma.Nome, ma.Profissao)
// }

// func (me MenorDeIdade) Estudar() string {
// 	return fmt.Sprintf("%s está no %d ano do colegial", me.Nome, me.AnoColegial)
// }

// func main() {

// 	ma := MaiorDeIdade{Pessoa{"João", 30}, "Programador"}
// 	me := MenorDeIdade{Pessoa{"Maria", 15}, 3}

// 	fmt.Println(ma.Trabalhar())
// 	fmt.Println(me.Estudar())

// }
