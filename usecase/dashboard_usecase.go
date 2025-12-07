package usecase

import (
	"fmt"

	"parry_end/model"
	"parry_end/repository"
)

type DashboardUsecase struct {
	curriculoRepository *repository.CurriculoRepository
	producaoRepository  *repository.ProducaoRepository
}

func NewDashboardUsecase(
	curriculorepo *repository.CurriculoRepository,
	producaorepo *repository.ProducaoRepository,
) DashboardUsecase {
	return DashboardUsecase{
		curriculoRepository: curriculorepo,
		producaoRepository:  producaorepo,
	}
}

func (ds *DashboardUsecase) GetRelatorioGeral() (*model.RelatorioGeral, error) {
	var relatorioGeral *model.RelatorioGeral = &model.RelatorioGeral{}
	var err error
	fmt.Println("Pegando Número de Curriculos:")
	relatorioGeral.TotalCurriculos, err = ds.curriculoRepository.GetCurriculoCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("Pegando Números de Produções:")
	relatorioGeral.TotalProducoes, err = ds.producaoRepository.GetProducaoCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("Pegando quantidade de Curricuoos Atualizados")
	relatorioGeral.CurriculosAtualizados, err = ds.curriculoRepository.GetUpdatedCurriculos()
	if err != nil {
		return nil, err
	}

	return relatorioGeral, nil
}

func (ds *DashboardUsecase) ConstructRelatorioAno(
	relatorioGeral *model.RelatorioGeral,
) error {
	fmt.Println("Pegando Producoes agrupadas por ano e tipo")
	relatorios, err := ds.producaoRepository.GetProducoesGroypByAnoTipo()
	if err != nil {
		return err
	}

	fmt.Println("Pegando numero de producoes por ano")
	producoesPorAno, err := ds.producaoRepository.GetProducoesCountByYear()

	var totalGeral int64 = 0
	for _, value := range producoesPorAno {
		fmt.Println("Construindo relatório geral")
		relatorios[value.Ano].QuantidadeDeContribuintes = value.Contagem

		totalGeral += relatorios[value.Ano].ProducoesTotal
	}

	relatorioGeral.Detalhes = relatorios

	relatorioGeral.TotalProducoes = &totalGeral

	return nil
}

//
// func (ds *DashboardUsecase)ConstructRelatorioAno(relatiorioGeral *model.RelatorioGeral)error{
//
// 	var relatorioAno []*model.RelatorioAno
//
// 	tuplasAnoTipo, err := ds.producaoRepository.GetProducoesYearAndType()
//
// 	if err != nil {
// 		return nil
// 	}
//
// 	for _,tupla := range tuplasAnoTipo {
//
// 	}

//}
