package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type PessoaUsecase struct {
	repository *repository.PessoaRepository
}

func NewPessoaUseCase(repo *repository.PessoaRepository) PessoaUsecase {
	return PessoaUsecase{
		repository: repo,
	}
}

func (pu *PessoaUsecase) GetPessoas() (*[]model.Pessoa, error) {
	return pu.repository.GetPessoas()
}

func (pu *PessoaUsecase) CreatePessoa(pessoa *model.Pessoa) error {
	err := pu.repository.CreatePessoa(pessoa)

	if err != nil {
		return err
	}

	return nil
}

func (pu *PessoaUsecase) GetPessoaByCPF(CPF int) (*model.Pessoa, error) {
	pessoa, err := pu.repository.GetPessoaByCPF(CPF)

	if err != nil {
		return nil, err
	}

	return pessoa, nil
}

func (pu *PessoaUsecase) UpdatePessoa(pessoa *model.Pessoa) error {
	err := pu.repository.UpdatePessoa(pessoa)

	if err != nil {
		return err
	}

	return nil
}
