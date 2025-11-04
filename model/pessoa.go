package model

type Pessoa struct {
	IdPessoa      int64  `json:"-"`
	Nome          string `json:"nome"`
	IdLattes      int    `json:"id_lattes"`
	Sexo          bool   `json:"sexo"`
	Abreviatura   string `json:"abreviatura"`
	Nacionalidade string `json:"nacionalidade"`
}
