package model

type Abreviatura struct {
	IdAbreviatura int64  `json:"-"`
	IdPessoa      *int64 `json:"-"` // O Ponteiro é necessário para passar NULL ao DB caso a abreviatura seja de um coautor e não de uma pessoa
	Abreviatura   string `json:"abreviatura"`
}
