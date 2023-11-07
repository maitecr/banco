package main

import (
	clientes "banco/cliente"
	contas "banco/conta"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	cliente01 := clientes.Titular{"Pafuncio", "123343", "dev"}

	conta01 := contas.ContaCorrente{
		Titular:       cliente01,
		NumeroAgencia: 123,
		NumeroConta:   234,
	}

	fmt.Println(conta01.Depositar(100))
	PagarBoleto(&conta01, 39)
	fmt.Println(conta01.GetSaldo())

	conta02 := contas.ContaPoupanca{}
	conta02.Depositar(300)
	//fmt.Println(conta02)
	PagarBoleto(&conta02, 7)
	fmt.Println(conta02.GetSaldo())

}
