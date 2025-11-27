package model

type Producao struct {
	IdProducao       int64          `json:"-"`
	Autor            string         `json:"autor"`
	Titulo           string         `json:"titulo"`
	DataDePublicacao string         `json:"data_de_publicacao"`
	TipoId           int64          `json:"tipo_id"`
	TipoS            string         `json:"tipo_s"`
	Hash             int64          `json:"hash"`
	Coautores        []*Abreviatura `json:"coautores"`
}

