package types

import (
	"net"
	"regexp"
	"strings"
)

type Email struct {  //type Email string
	Email string
}

func (e *Email) NewEmail(email string) Email {

	if e.validarEmail(email) != true {
		return Email{}
	}

	return Email{
		Email: email,
	}

}

func (e *Email) validarEmail(email string) bool {
	// Verificação básica de comprimento
	if len(email) < 6 || len(email) > 254 {
		return false
	}

	// Verifica se contém @ e posição válida
	atPos := strings.Index(email, "@")
	if atPos < 1 || atPos == len(email)-1 {
		return false
	}

	localPart := email[:atPos]
	domainPart := email[atPos+1:]

	// Valida parte local (antes do @)
	if !e.validarParteLocal(localPart) {
		return false
	}

	// Valida domínio (depois do @)
	if !e.validarDominio(domainPart) {
		return false
	}

	return true
}

func (e *Email) validarParteLocal(localPart string) bool {
	// Regex para parte local (mais permissiva)
	// Permite: letras, números, ! # $ % & ' * + - / = ? ^ _ ` { | } ~
	localRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+$`
	matched, _ := regexp.MatchString(localRegex, localPart)
	if !matched {
		return false
	}

	// Não pode começar ou terminar com ponto
	if strings.HasPrefix(localPart, ".") || strings.HasSuffix(localPart, ".") {
		return false
	}

	// Não pode ter dois pontos consecutivos
	if strings.Contains(localPart, "..") {
		return false
	}

	// Comprimento máximo da parte local é 64 caracteres
	if len(localPart) > 64 {
		return false
	}

	return true
}

func (e *Email) validarDominio(domain string) bool {
	// Verifica domínio IP entre colchetes
	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		ip := domain[1 : len(domain)-1]
		return e.validarIP(ip)
	}

	// Domínio regular
	return e.validarDominioRegular(domain)
}

func (e *Email) validarDominioRegular(domain string) bool {
	// Regex para domínio (mais estrita)
	domainRegex := `^[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?)*$`
	matched, _ := regexp.MatchString(domainRegex, domain)
	if !matched {
		return false
	}

	// Divide o domínio em partes
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return false // Precisa ter pelo menos domínio e TLD
	}

	// Verifica cada parte do domínio
	for _, part := range parts {
		// Cada parte não pode ser vazia
		if len(part) == 0 {
			return false
		}

		// Cada parte não pode começar ou terminar com hífen
		if strings.HasPrefix(part, "-") || strings.HasSuffix(part, "-") {
			return false
		}

		// Comprimento máximo de cada parte é 63 caracteres
		if len(part) > 63 {
			return false
		}
	}

	// Verifica TLD (última parte)
	tld := parts[len(parts)-1]
	if len(tld) < 2 {
		return false // TLD muito curto
	}

	// TLD deve conter apenas letras
	tldRegex := `^[a-zA-Z]+$`
	matched, _ = regexp.MatchString(tldRegex, tld)
	if !matched {
		return false
	}

	return true
}

func (e *Email) validarIP(ip string) bool {
	// Verifica se é um endereço IP válido
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

// Versão alternativa usando regex única (mais simples)
func (e *Email) validarEmailSimples(email string) bool {
	// Regex mais comum para validação de email
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

	matched, _ := regexp.MatchString(emailRegex, email)
	if !matched {
		return false
	}

	// Verificações adicionais
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	localPart := parts[0]
	domainPart := parts[1]

	// Verifica comprimentos
	if len(localPart) > 64 || len(domainPart) > 253 {
		return false
	}

	return true
}

func (e *Email) GetEmail() string {
	return e.Email
}

func (e *Email) SetEmail(email string) bool {
	if e.validarEmail(email) {
		e.Email = email
		return true
	}

	return false
}
