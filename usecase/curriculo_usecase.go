package usecase

import (
	"database/sql"
	"fmt"

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
	fmt.Println("Pegando Curriculo:")
	return cu.curriculoRepository.GetCurriculos()
}

func (cu *CurriculoUsecase) identifyTipoProducao(
	producao *model.Producao,
) int64 {
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
	fmt.Println("Pegando Curriculo Por Id:", idPessoa)
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

		fmt.Println("Pegando Producao por Hash:", values)
		_, err := cu.producaoRepository.GetProducaoByHash(values)

		if err == sql.ErrNoRows {

			fmt.Println("CreateProducao", values, curriculo)
			cu.producaoRepository.CreateProducao(values, curriculo)
		}

	}

	fmt.Println("Update Curriuclo", curriculo)
	err := cu.curriculoRepository.UpdateCurriculo(curriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteProducaoByIdCurriculo(
	idCurriculo int64,
) error {
	fmt.Println("Deletando Producao Por Id Curriculo", idCurriculo)
	err := cu.producaoRepository.DeleteProducaoByIdCurriculo(idCurriculo)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteCurriculoByIdPessoa(idPessoa int64) error {
	fmt.Println("Deletando Curriculo por Id Pessoa", idPessoa)
	err := cu.curriculoRepository.DeleteCurriculo(idPessoa)
	if err != nil {
		return err
	}

	return nil
}

func (cu *CurriculoUsecase) DeleteCoautoresByIdProducao(
	idProducao int64,
) error {
	fmt.Println("Deletando Coautor Por id Producao", idProducao)
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
