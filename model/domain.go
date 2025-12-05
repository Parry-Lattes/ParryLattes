package model

type PessoaCurriculo struct {
	Pessoa    *Pessoa
	Curriculo *Curriculo
}

type Pessoa struct {
	IdPessoa      int64          `json:"-"`
	Nome          string         `json:"nome"`
	IdLattes      string         `json:"id_lattes"`
	Abreviaturas  []*Abreviatura `json:"abreviaturas"`
	Nacionalidade string         `json:"nacionalidade"`
}

type Curriculo struct {
	IdCurriculo       int64       `json:"-"`
	IdPessoa          int         `json:"-"`
	UltimaAtualizacao string      `json:"ultima_atualizacao"`
	Producoes         []*Producao `json:"producoes"`
}

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

type Coautor struct {
	IdCoautor   int64        `json:"-"`
	IdProducao  int64        `json:"-"`
	Abreviatura *Abreviatura `json:"abreviatura"`
}

type Abreviatura struct {
	IdAbreviatura int64  `json:"-"`
	IdPessoa      *int64 `json:"-"` // O Ponteiro é necessário para passar NULL ao DB caso a abreviatura seja de um coautor e não de uma pessoa
	Abreviatura   string `json:"abreviatura"`
}

type Login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
