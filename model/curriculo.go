package model

type Curriculo struct {
	IdLattes          string     `json:"id_lattes"`
	UltimaAtualizacao string     `json:"ultima_atualizacao"`
	Producoes         []Producao `json:"producoes"`
}
