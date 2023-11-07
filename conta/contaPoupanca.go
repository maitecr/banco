package contas

import clientes "banco/cliente"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso. saldo:", c.saldo
	} else {
		return "Valor inválido para depósito saldo:", c.saldo
	}
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
