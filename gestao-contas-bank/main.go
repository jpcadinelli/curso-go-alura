package main

import (
	"fmt"
	"gestaoContasBank/clientes"
	"gestaoContasBank/contas"
)

func main() {
	contaDoJoao := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "João Pedro",
		CPF:       "123.456.789-00",
		Profissao: "Engenheiro de Software"},
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 158.99}
	fmt.Println(contaDoJoao)
}
