package model

type Producao struct {
	IdProducao        int64  `json:"-"`
	Autor             string `json:"autor"`
	Titulo            string `json:"titulo"`
	Descricao         string `json:"descricao"`
	Link              string `json:"link"`
	DataDePublicacao  string `json:"data_de_publicacao"`
	Tipo              int64  `json:"tipo"`
	IdGrupoDePesquisa int64  `json:"id_grupo_de_pesquisa"`
}
