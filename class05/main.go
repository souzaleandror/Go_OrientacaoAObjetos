package main

import (
	"fmt"
	"class05/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
} 

type verificarConta interface {
	Sacar(valorDoBoleto float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}
	fmt.Println(contaDaPati)

	contaDoDenis.Depositar(-100)

	fmt.Println(contaDoDenis)
	fmt.Println(contaDoDenis.ObterSaldo())
	contaDoDenis.Depositar(5000)
	contaDaPati.Depositar(3000)
	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(contaDoDenis.ObterSaldo())
	contaDoDenis.Sacar(500)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 500)
	PagarBoleto(&contaDaPati, 500)
	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(contaDaPati.ObterSaldo())
}