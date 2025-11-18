package usecase

import (
	"database/sql"
	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	CurriculoRepository *repository.CurriculoRepository
	ProducaoRepository  *repository.ProducaoRepository
}

func NewCurriculoUseCase(curriculorepo *repository.CurriculoRepository, producaorepo *repository.ProducaoRepository) CurriculoUsecase {
	return CurriculoUsecase{
		CurriculoRepository: curriculorepo,
		ProducaoRepository:  producaorepo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() ([]*model.Curriculo, error) {
	return cu.CurriculoRepository.GetCurriculos()
}

func (cu *CurriculoUsecase) GetCurriculoById(idPessoa int) (*model.Curriculo, error) {

	curriculo, err := cu.CurriculoRepository.GetCurriculoById(idPessoa)

	if err != nil {
		return nil, err
	}

	curriculo.Producoes, err = cu.ProducaoRepository.GetProducaoByIdLattes(curriculo)

	if err != nil {
		return nil, err
	}

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

func (cu *CurriculoUsecase) DeleteProducao(producao model.Producao) error {
	err := cu.ProducaoRepository.DeleteProducao(producao.Hash)

	if err != nil {
		return err
	}

	return nil
}
