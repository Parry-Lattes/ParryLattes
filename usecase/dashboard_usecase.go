package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type DashboardUsecase struct {
	CurriculoRepository *repository.CurriculoRepository
	ProducaoRepository  *repository.ProducaoRepository
}

func NewDashboardUsecase(
	curriculorepo *repository.CurriculoRepository,
	producaorepo *repository.ProducaoRepository,
) DashboardUsecase {
	return DashboardUsecase{
		CurriculoRepository: curriculorepo,
		ProducaoRepository:  producaorepo,
	}
}

func (ds *DashboardUsecase) GetRelatorioGeral() (*model.RelatorioGeral, error) {
	var relatorioGeral *model.RelatorioGeral = &model.RelatorioGeral{}
	var err error
	relatorioGeral.TotalCurriculos, err = ds.CurriculoRepository.GetCurriculoCount()
	if err != nil {
		return nil, err
	}

	relatorioGeral.TotalProducoes, err = ds.ProducaoRepository.GetProducaoCount()
	if err != nil {
		return nil, err
	}

	return relatorioGeral, nil
}

func (ds *DashboardUsecase) ConstructRelatorioAno(
	relatorioGeral *model.RelatorioGeral,
) error {
	relatorios, err := ds.ProducaoRepository.GetProducoesGroypByAnoTipo()
	if err != nil {
		return err
	}

	relatorioGeral.Detalhes = relatorios

	// Calcula o total geral
	var totalGeral int64 = 0
	for _, relatorio := range relatorios {
		if relatorio.ProducoesTotal != nil {
			totalGeral += *relatorio.ProducoesTotal
		}
	}

	relatorioGeral.TotalProducoes = &totalGeral

	return nil
}

//
// func (ds *DashboardUsecase)ConstructRelatorioAno(relatiorioGeral *model.RelatorioGeral)error{
//
// 	var relatorioAno []*model.RelatorioAno
//
// 	tuplasAnoTipo, err := ds.ProducaoRepository.GetProducoesYearAndType()
//
// 	if err != nil {
// 		return nil
// 	}
//
// 	for _,tupla := range tuplasAnoTipo {
//
// 	}

//}
