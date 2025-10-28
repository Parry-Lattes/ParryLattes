package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type ProducaoUsecase struct {
	Repository repository.ProducaoRepository
}

func NewProducaoUseCase(repo repository.ProducaoRepository) ProducaoUsecase {
	return ProducaoUsecase{
		Repository: repo,
	}
}

func (pu *ProducaoUsecase) GetProducoes() (*[]model.Producao, error) {
	return pu.Repository.GetProducoes()
}

func (pu *ProducaoUsecase) GetProducaoById(idProducao int) (*[]model.Producao, error) {
	producao, err := pu.Repository.GetProducaoById(idProducao)

	if err != nil {
		return nil, err
	}

	return producao, nil
}

// func (pu *ProducaoUsecase) CreateProducao(producao model.Producao,idLattes int)
