package main

import (
	"fmt"

	// "github.com/AndreCDiniz/banco/clientes"
	"github.com/AndreCDiniz/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaMary := contas.ContaCorrente{}
	contaDaMary.Depositar(200)
	PagarBoleto(&contaDaMary, 100)

	fmt.Println(contaDaMary.ObterSaldo())
}
