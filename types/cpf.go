package types

import (
	"strconv"
	"strings"
)

type CPF struct {
	CPF string
}

func (c *CPF) NewCPF(cpf string) CPF {

	if c.validadeCPF(cpf) {
		return CPF{}
	}
	return CPF{
		CPF: cpf,
	}
}

func (c *CPF) validadeCPF(cpf string) bool {
	// Remove caracteres não numéricos
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifica se tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais (CPF inválido)
	if c.todosDigitosIguais(cpf) {
		return false
	}

	// Calcula primeiro dígito verificador
	digito1 := c.calcularDigito(cpf, 10)

	// Calcula segundo dígito verificador
	digito2 := c.calcularDigito(cpf, 11)

	// Verifica se os dígitos calculados batem com os informados
	digito1Informado, _ := strconv.Atoi(string(cpf[9]))
	digito2Informado, _ := strconv.Atoi(string(cpf[10]))

	return digito1 == digito1Informado && digito2 == digito2Informado
}

func (c *CPF) calcularDigito(cpf string, pesoInicial int) int {
	soma := 0
	peso := pesoInicial

	// Calcula a soma ponderada dos dígitos
	for i := 0; i < pesoInicial-1; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma += digito * peso
		peso--
	}

	// Calcula o resto da divisão por 11
	resto := soma % 11

	// Define o dígito verificador
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func (c *CPF) todosDigitosIguais(cpf string) bool {
	primeiroDigito := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != primeiroDigito {
			return false
		}
	}
	return true
}

func (c *CPF) GetCPF() string {
	return c.CPF
}

func (c *CPF) SetCPF(cpf string) bool {

	if c.validadeCPF(cpf) {
		c.CPF = cpf
		return true
	}

	return false

}
