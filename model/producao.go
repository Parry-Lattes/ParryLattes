package model

type Producao struct {
	IdProducao       int64      `json:"-"`
	Autor            string     `json:"autor"`
	Titulo           string     `json:"titulo"`
	DataDePublicacao string     `json:"data_de_publicacao"`
	TipoId           int64      `json:"-"`
	Tipo             string     `json:"tipo"`
	Hash             string     `json:"hash"`
	Coautores        []*Coautor `json:"coautores"`
}
