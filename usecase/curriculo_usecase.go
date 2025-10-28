package usecase

import (
	"fmt"
	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	Repository repository.CurriculoRepository
}

func NewCurriculoUseCase(repo repository.CurriculoRepository) CurriculoUsecase {
	return CurriculoUsecase{
		Repository: repo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() (*[]model.Curriculo, error) {
	return cu.Repository.GetCurriculos()
}

func (cu *CurriculoUsecase) GetCurriculoById(idLattes int) (*model.Curriculo, error) {
	curriculo, err := cu.Repository.GetCurriculoById(idLattes)
	
	if err != nil {
		return nil, err
	}

	return curriculo, nil
}

func (cu *CurriculoUsecase) CreateCurriculo(curriculo *model.Curriculo, idPessoa int) (*model.Curriculo, error) {
	IdCurriculo, err := cu.Repository.CreateCurriculo(curriculo, idPessoa)

	if err != nil {
		return &model.Curriculo{}, err
	}

	fmt.Println(IdCurriculo)

	return curriculo, nil
}

func (cu *CurriculoUsecase) UpdateCurriculo(curriculo *model.Curriculo) error {
	err := cu.Repository.UpdateCurriculo(curriculo)

	if err != nil {
		return err
	}

	return nil
}
