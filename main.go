package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/* realizar uma requisicao http e salvar o conteudo num arquivo de texto*/
func main() {

	fmt.Println("iniciando execução...")
	sites := leArquivoDeSites()
	salvaConteudoHTTP(sites)
	fmt.Println("fim execução.")
}

func leArquivoDeSites() []string {
	var conteudo []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		conteudo = append(conteudo, linha)

		if err != nil {
			fmt.Println("Erro ao ler linha: ", err)
			break
		}
	}
	defer arquivo.Close()

	return conteudo
}

func salvaConteudoHTTP(sites []string) {
	conteudo, err := os.OpenFile("requests.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir/criar arquivo", err)
	}

	for _, site := range sites {
		fmt.Printf(site)

		result, _ := http.Get(site)

		defer result.Body.Close()
		body, _ := io.ReadAll(result.Body)

		conteudo.WriteString(string(body))
	}

	defer conteudo.Close()
}
