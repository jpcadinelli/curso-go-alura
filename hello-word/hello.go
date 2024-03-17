package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	fmt.Println("")
	exibeIntrodução()
	fmt.Println("")
	for {
		exibeMenu()
		fmt.Println("")
		comando := leComando()
		fmt.Println("")
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}
		fmt.Println("")
	}
}

func exibeIntrodução() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://www.caelum.com.br", "https://www.kabum.com.br"}
	for i := 0; i < monitoramentos; i++ {
		for i := 0; i < len(sites); i++ {
			resp, _ := http.Get(sites[i])
			if resp.StatusCode == 200 {
				fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
			} else {
				fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
			}
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}
