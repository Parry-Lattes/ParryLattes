package model

type Producao struct {
	IdProducao       int64  `json:"-"`
	Autor            string `json:"autor"`
	Titulo           string `json:"titulo"`
	Descricao        string `json:"descricao"`
	Link             string `json:"link"`
	DataDePublicacao string `json:"data_de_publicacao"`
	TipoId           int64  `json:"tipo_id"`
	TipoS            string `json:"tipo_s"`
	Hash             int64    `json:"hash"`
}
