package model

type Coautor struct {
	IdCoautor   int64        `json:"-"`
	IdProducao  int64        `json:"-"`
	Abreviatura *Abreviatura `json:"abreviatura"`
}
