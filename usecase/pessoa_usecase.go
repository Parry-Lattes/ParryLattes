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

func (pu *PessoaUsecase) GetPessoas() ([]*model.Pessoa, error) {

	pessoas,err := pu.repository.GetPessoas()
	
	if err != nil{
		return nil,err
	}
	
	for _,values := range pessoas{
		values.Abreviaturas,err = pu.repository.GetAbreviaturasById(values.IdPessoa)				
	}

	return pessoas, nil
}

func (pu *PessoaUsecase) CreatePessoa(pessoa *model.Pessoa) error {
	pessoa,err := pu.repository.CreatePessoa(pessoa)

	if err != nil {
		return err
	}

	for _, value := range pessoa.Abreviaturas {
		
		value.IdPessoa = pessoa.IdPessoa
		err = pu.repository.CreateAbreviatura(value)
	
		if err != nil {
			return err
		}
	}
	 
	return nil
}

func (pu *PessoaUsecase) GetPessoaByIdLattes(idLattes int) (*model.Pessoa, error) {
	
	pessoa, err := pu.repository.GetPessoaByIdLattes(idLattes)

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

// func (pu)
