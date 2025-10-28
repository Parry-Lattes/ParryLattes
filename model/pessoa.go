package model

type Pessoa struct {
	Nome          string `json:"nome"`
	CPF           int `json:"cpf"`
	Sexo          bool   `json:"sexo"`
	Abreviatura   string `json:"abreviatura"`
	Nacionalidade string `json:"nacionalidade"`
}
