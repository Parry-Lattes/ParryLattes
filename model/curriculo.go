package model

type Curriculo struct {
	IdCurriculo       int64       `json:"-"`
	IdPessoa          int         `json:"-"`
	UltimaAtualizacao string      `json:"ultima_atualizacao"`
	Producoes         []*Producao `json:"producoes"`
}
