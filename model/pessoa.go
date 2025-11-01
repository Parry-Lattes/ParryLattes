package model

type Pessoa struct {
	IdPessoa      int64  `json:"-"`
	Nome          string `json:"nome"`
	CPF           int    `json:"cpf"`
	Sexo          bool   `json:"sexo"`
	Abreviatura   string `json:"abreviatura"`
	Nacionalidade string `json:"nacionalidade"`
}
