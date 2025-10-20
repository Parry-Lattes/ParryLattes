package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type PessoaUsecase struct {
	repository repository.PessoaRepository
}

func NewPessoaUseCase(repo repository.PessoaRepository) PessoaUsecase {
	return PessoaUsecase{
		repository: repo,
	}
}

func (pu *PessoaUsecase) GetPessoas() ([]model.Pessoa, error) {
	return pu.repository.GetPessoas()
}
