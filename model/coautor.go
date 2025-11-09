package model

type Coautor struct {
	IdCoautor  int64  `json:"id_coautor"`
	IdProducao int64  `json:"id_producao"`
	Coautor    string `json:"coautor"`
}
