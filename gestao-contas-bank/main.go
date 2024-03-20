package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}


func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso.", c.saldo
	} else {
		return "Saldo insuficiente.", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	} else {
		return "Valor negativo não pode ser depositado.", c.saldo
	}
}

func main() {
	contaDoJoao := ContaCorrente{titular: "João Pedro", numeroAgencia: 589, numeroConta: 123456, saldo: 158.99}
	fmt.Println(contaDoJoao.saldo)
	fmt.Println(contaDoJoao.Depositar(100))
}
