package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
)

const monitoring = 5
const delay = 5

func main() {
	showIntroduction()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo...")
			showLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")

			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Thiago"
	version := 1.1

	fmt.Println("Olá", name)
	fmt.Println("O programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)

	return readedCommand
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://cobogo-strapi.herokuapp.com", "https://app.cobogo.social", "https://cobogo.social"}

	sites := readSitesFile()

	for i := 0; i < monitoring; i++ {
		fmt.Println("Carregando...")

		for _, site := range sites {
			testSite(site)

		}

		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		createLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		createLogs(site, false)
	}
}

func readSitesFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func createLogs(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
