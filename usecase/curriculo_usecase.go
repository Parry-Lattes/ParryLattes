package usecase

import (
	"fmt"
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

func (cu *CurriculoUsecase) GetCurriculos() (*[]model.Curriculo, error) {
	return cu.repository.GetCurriculos()
}

func (cu *CurriculoUsecase) GetCurriculoById(idCurriculo int) (*model.Curriculo, error) {
	curriculo, err := cu.repository.GetCurriculoById(idCurriculo)

	if err != nil {
		return nil, err
	}

	return curriculo, nil
}

func (cu *CurriculoUsecase) CreateCurriculo(curriculo *model.Curriculo, idPessoa int) (*model.Curriculo, error) {
	IdCurriculo, err := cu.repository.CreateCurriculo(curriculo, idPessoa)

	if err != nil {
		return &model.Curriculo{}, err
	}

	fmt.Println(IdCurriculo)

	return curriculo, nil
}

func (cu *CurriculoUsecase) UpdateCurriculo(curriculo *model.Curriculo) error {
	err := cu.repository.UpdateCurriculo(curriculo)

	if err != nil {
		return err
	}

	return nil
}
