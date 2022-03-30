package main

import "fmt"

type ContaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", agencia: 589, numeroConta: 123456, saldo: 120.45}
	contaDaVal := ContaCorrente{"Val", 222, 111222, 200.45}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaVal)
}
