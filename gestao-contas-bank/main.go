package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}


func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

func main() {
	contaDoJoao := ContaCorrente{titular: "João Pedro", numeroAgencia: 589, numeroConta: 123456, saldo: 158.99}
	fmt.Println(contaDoJoao.saldo)
	fmt.Println(contaDoJoao.Sacar(-100))
	fmt.Println(contaDoJoao.saldo)
}
