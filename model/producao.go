package model

type Producao struct {
	Autor            string `json:"autor"`
	Titulo           string `json:"titulo"`
	Descricao        string `json:"descricao"`
	Link             string `json:"link"`
	DataDePublicacao string `json:"data_de_publicacao"`
	Tipo             string `json:"tipo"`
}
