package model

type RelatorioGeral struct {
	TotalCurriculos int64                `json:"total_curriculos"`
	TotalProducoes  int64                `json:"total_producoes"`
	Detalhes        map[int]RelatorioAno `json:"detalhes" `
}

type RelatorioAno struct {
	QuantidadeDeContribuintes int64 `json:"qtd_contribuintes"`
	ProducoesBibliograficas   int64 `json:"Bibliográficas"`
	ProducoesTecnicas         int64 `json:"Técnica"`
	ProducoesPatente          int64 `json:"Patente"`
	ProducoesOutro            int64 `json:"Outro"`
	ProducoesTotal            int64 `json:"total_producoes"`
}
