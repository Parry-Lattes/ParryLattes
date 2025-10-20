package model

type Pessoa struct {
	IdPessoa      int64  `json:"id_pessoa"`
	Nome          string `json:"nome"`
	Sexo          bool   `json:"sexo"`
	Abreviatura   string `json:"abreviatura"`
	Nacionalidade string `json:"nacionalidade"`
}
