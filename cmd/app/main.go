// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	//"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	//"github.com/seu-usuario/meu-projeto-go/internal/hello"

	calculoimc "github.com/seu-usuario/meu-projeto-go/internal/calculoIMC"
)

// Função principal do programa
func main() {
	// Mensagem inicial da aplicação
	//fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamada para função de saudação
	//hello.SayHello()
	fmt.Println("Calcule seu Imc!!")
	var Altura float64
	var Peso float64
	fmt.Printf("digite Sua Altura:")
	fmt.Scanln(&Altura)
	fmt.Printf("Digite seu peso:")
	fmt.Scanln(&Peso)

	imc := calculoimc.CalculoIMC(Peso, Altura)
	fmt.Println("Altura digitada", Altura, "Peso digitado:", Peso)
	fmt.Printf("Seu IMC é: %.2f", imc)
	fmt.Println()
	if imc < 18.5 {
		fmt.Println("Baixo Peso")
	} else if imc >= 18.5 && imc < 24.9 {
		fmt.Println("Peso Normal")
	} else if imc >= 25 && imc < 29.9 {
		fmt.Println("Sobrepeso")
	} else if imc >= 30 && imc < 34.9 {
		fmt.Println("Obesidade Grau 1")
	} else if imc >= 35 && imc < 39.9 {
		fmt.Println("Obesidade Grau 2")
	} else if imc >= 40 {
		fmt.Println("Obesidade Grau 3")
	}

	// Demonstração: cálculo do 10º número de Fibonacci
	//n := 10
	// Chama a função Fibonacci do pacote fibonacci
	// fibonacci // importado acima
	// Fibonacci(n) // retorna o n-ésimo número da sequência
	// := é usado para declarar e inicializar a variável
	//valor := fibonacci.Fibonacci(n)
	// Imprime o resultado com formatação
	//fmt.Printf("F(%d) = %d\n", n, valor)

	// Demonstração: imprimir a sequência completa até n
	//fibonacci.PrintSequence(n)
}
