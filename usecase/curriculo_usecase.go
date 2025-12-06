package usecase

import (
	"database/sql"

	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	curriculoRepository   *repository.CurriculoRepository
	producaoRepository    *repository.ProducaoRepository
	abreviaturaRepository *repository.AbreviaturaRepository
}

func NewCurriculoUseCase(
	curriculorepo *repository.CurriculoRepository,
	producaorepo *repository.ProducaoRepository,
	abreviaturarepo *repository.AbreviaturaRepository,
) CurriculoUsecase {
	return CurriculoUsecase{
		curriculoRepository:   curriculorepo,
		producaoRepository:    producaorepo,
		abreviaturaRepository: abreviaturarepo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() ([]*model.Curriculo, error) {
	return cu.curriculoRepository.GetCurriculos()
}

func (cu *CurriculoUsecase) identifyTipoProducao(producao *model.Producao) int64 {
	switch producao.Tipo {
	case "Bibliográfica":
		return 1
	case "Técnica":
		return 2
	case "Patente":
		return 3
	default:
		return 4
	}
}

func (cu *CurriculoUsecase) GetCurriculoById(
	idPessoa int64,
) (*model.Curriculo, error) {
	curriculo, err := cu.curriculoRepository.GetCurriculoById(idPessoa)
	if err != nil {
		return nil, err
	}

	// curriculo.Producoes, err = cu.producaoRepository.GetProducaoByIdLattes(
	// 	curriculo,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	//
	// for _, producao := range curriculo.Producoes {
	// 	producao.Coautores, err = cu.abreviaturaRepository.GetCoautoresByIdProducao(
	// 		producao.IdProducao,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	for _, coautor := range producao.Coautores {
	// 		coautor.Abreviatura, err = cu.abreviaturaRepository.GetAbreviaturaByCoautor(
	// 			coautor,
	// 		)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	//
	// 	}
	//
	// 	// }

	return curriculo, nil
}

func (cu *CurriculoUsecase) UpdateCurriculo(curriculo *model.Curriculo) error {
	for _, values := range curriculo.Producoes {

		_, err := cu.producaoRepository.GetProducaoByHash(values)

		if err == sql.ErrNoRows {
			cu.producaoRepository.CreateProducao(values, curriculo)
		}

	}
	err := cu.curriculoRepository.UpdateCurriculo(curriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteProducaoByIdCurriculo(
	idCurriculo int64,
) error {
	err := cu.producaoRepository.DeleteProducaoByIdCurriculo(idCurriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteCurriculoByIdPessoa(idPessoa int64) error {
	err := cu.curriculoRepository.DeleteCurriculo(idPessoa)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteCoautoresByIdProducao(
	idProducao int64,
) error {
	err := cu.abreviaturaRepository.DeleteCoautoresByIdProducao(idProducao)
	if err != nil {
		return err
	}

	return nil
}

// func (cu *curriculoUsecase) DeleteProducao(producao model.Producao) error {
// 	err := cu.producaoRepository.DeleteProducao(producao.Hash)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
