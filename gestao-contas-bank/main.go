package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJoao := ContaCorrente{titular: "João Pedro", numeroAgencia: 589, numeroConta: 123456, saldo: 158.99}
	contaDaBruna := ContaCorrente{titular: "Bruna", numeroAgencia: 589, numeroConta: 654321, saldo: 1345.00}
	fmt.Println(contaDoJoao)
	fmt.Println(contaDaBruna)
}
