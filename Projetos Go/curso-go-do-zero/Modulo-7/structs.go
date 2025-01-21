// métodos em struct

package main

// import "fmt"

// type Pessoa struct {
// 	Nome      string
// 	Sobrenome string
// 	Idade     int
// 	CPF       string
// }

// type Funcionario struct {
// 	Pessoa
// 	Cargo   string
// 	Salario float32
// }

// func main() {
// 	p1 := Pessoa{"João", "Silva", 30, "123.456.789-00"}
// 	f1 := Funcionario{Pessoa{"Maria", "Silva", 25, "987.654.321-00"}, "Gerente", 5250.25}

// 	fmt.Println(p1.ListaNomeCompletoEIdade())
// 	fmt.Println(f1.ListaFuncionario())
// }

// func (p Pessoa) ListaNomeCompletoEIdade() string {
// 	return fmt.Sprintf("%s %s tem %d anos", p.Nome, p.Sobrenome, p.Idade)
// }

// func (f Funcionario) ListaFuncionario() string {
// 	return fmt.Sprintf("%s %s tem o cargo: %s e recebe R$%f por mês", f.Nome, f.Sobrenome, f.Cargo, f.Salario)
// }
