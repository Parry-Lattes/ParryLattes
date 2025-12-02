package usecase

import (
	"database/sql"

	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	CurriculoRepository   *repository.CurriculoRepository
	ProducaoRepository    *repository.ProducaoRepository
	AbreviaturaRepository *repository.AbreviaturaRepository
}

func NewCurriculoUseCase(
	curriculorepo *repository.CurriculoRepository,
	producaorepo *repository.ProducaoRepository,
	abreviaturarepo *repository.AbreviaturaRepository,
) CurriculoUsecase {
	return CurriculoUsecase{
		CurriculoRepository:   curriculorepo,
		ProducaoRepository:    producaorepo,
		AbreviaturaRepository: abreviaturarepo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() ([]*model.Curriculo, error) {
	return cu.CurriculoRepository.GetCurriculos()
}

func (cu *CurriculoUsecase) identifyTipo(producao *model.Producao) int64 {
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
	curriculo, err := cu.CurriculoRepository.GetCurriculoById(idPessoa)
	if err != nil {
		return nil, err
	}

	// curriculo.Producoes, err = cu.ProducaoRepository.GetProducaoByIdLattes(
	// 	curriculo,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	//
	// for _, producao := range curriculo.Producoes {
	// 	producao.Coautores, err = cu.AbreviaturaRepository.GetCoautoresByIdProducao(
	// 		producao.IdProducao,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	for _, coautor := range producao.Coautores {
	// 		coautor.Abreviatura, err = cu.AbreviaturaRepository.GetAbreviaturaByCoautor(
	// 			coautor,
	// 		)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	//
	// 	}
	//
	// }

	return curriculo, nil
}

func (cu *CurriculoUsecase) UpdateCurriculo(curriculo *model.Curriculo) error {
	for _, values := range curriculo.Producoes {

		_, err := cu.ProducaoRepository.GetProducaoByHash(values)

		if err == sql.ErrNoRows {
			cu.ProducaoRepository.CreateProducao(values, curriculo)
		}

	}
	err := cu.CurriculoRepository.UpdateCurriculo(curriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteProducaoByIdCurriculo(
	idCurriculo int64,
) error {
	err := cu.ProducaoRepository.DeleteProducaoByIdCurriculo(idCurriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteCurriculoByIdPessoa(idPessoa int64) error {
	err := cu.CurriculoRepository.DeleteCurriculo(idPessoa)
	if err != nil {
		return err
	}

	return nil
}

// func (cu *CurriculoUsecase) DeleteProducao(producao model.Producao) error {
// 	err := cu.ProducaoRepository.DeleteProducao(producao.Hash)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
