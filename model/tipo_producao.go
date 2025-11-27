package model

type _TipoProducao struct {
	Nome string
	ID int
}

type TipoProducao _TipoProducao

var (
	_ TipoProducao
	Bibliografica = TipoProducao{Nome: "Bibliográfica", ID: 0}
	Tecnica = TipoProducao{Nome: "Técnica", ID: 1}
	Patente = TipoProducao{Nome: "Patente", ID: 2}
	Outro = TipoProducao{Nome: "Outro", ID: 3}
)

func NewTipoProducao()
