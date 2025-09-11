// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	fmt.Println("Minha Primeira alteração em Go!!!")
	hello.SayHello()
	//Funçao para retorno Sayhello.

	retorno := fibonacci.Fibonacci(10)
	fmt.Println("Resultado fibifibonacci", retorno)
	//Funçao Fibonacci que retorna resultados
}
