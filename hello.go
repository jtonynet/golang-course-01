package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const testeEsperaEmSeg = 3

func main() {

	leSitesDoArquivo()

	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs")
		case 3:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco esse comando")
		}
	}

}

func exibeIntroducao() {
	versao := 1.2
	nome := "Epitombas"

	fmt.Println("Olar, Sr.", nome)
	fmt.Println("Versao do programa:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("3- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			site := site
			testaSite(site)
		}
		time.Sleep(testeEsperaEmSeg * time.Second)
	}
}

func testaSite(site string) {
	var successStatus int = 200
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == successStatus {
		fmt.Println("O site ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site ", site, "deu ruim!")
	}
}

func leSitesDoArquivo() []string {
	//arquivo, err := os.Open("sites.txt")
	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))
	return []string{}
}
