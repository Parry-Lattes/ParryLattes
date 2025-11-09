package model

type Abreviatura struct {
	IdAbreviatura int64  `json:"-"`
	IdPessoa      int64  `json:"-"`
	Abreviatura   string `json:"abreviatura"`
}
