package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const testeEsperaEmSeg = 3

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs")
			imprimeLogs()
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
		registraLog(site, true)
	} else {
		fmt.Println("O site ", site, "deu ruim!")
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Ocorreu um erro", err)
		}

		sites = append(sites, strings.TrimSpace(linha))
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
