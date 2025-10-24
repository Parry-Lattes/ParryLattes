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

func (pu *PessoaUsecase) GetPessoas() (*[]model.Pessoa, error) {
	return pu.repository.GetPessoas()
}

func (pu *PessoaUsecase) CreatePessoa(pessoa *model.Pessoa) (*model.Pessoa, error) {
	IdPessoa, err := pu.repository.CreatePessoa(pessoa)

	if err != nil {
		return &model.Pessoa{}, err
	}

	pessoa.IdPessoa = int64(IdPessoa)

	return pessoa, nil
}

func (pu *PessoaUsecase) GetPessoaById(idPessoa int) (*model.Pessoa, error) {
	pessoa, err := pu.repository.GetPessoaById(idPessoa)

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
