package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	repository repository.CurriculoRepository
}

func NewCurriculoUseCase(repo repository.CurriculoRepository) CurriculoUsecase {
	return CurriculoUsecase{
		repository: repo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() ([]model.Curriculo, error) {
	return cu.repository.GetCurriculos()
}

func (cu *CurriculoUsecase) GetCurriculoById(idCurriculo int) (*model.Curriculo, error) {
	curriculo, err := cu.repository.GetCurriculoById(idCurriculo)

	if err != nil {
		return nil, err
	}

	return curriculo, nil
}
