package model

type Pessoa struct {
	IdPessoa      int64          `json:"-"`
	Nome          string         `json:"nome"`
	IdLattes      int            `json:"id_lattes"`
	Abreviaturas  []*Abreviatura `json:"abreviaturas"`
	Nacionalidade string         `json:"nacionalidade"`
}
