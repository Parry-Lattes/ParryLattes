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
	return pu.GetProducoes()
}
