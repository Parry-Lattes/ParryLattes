package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type PessoaUsecase struct {
	pessoaRepository      *repository.PessoaRepository
	abreviaturaRepository *repository.AbreviaturaRepository
}

func NewPessoaUseCase(pessoarepo *repository.PessoaRepository, abreviaturarepo *repository.AbreviaturaRepository) PessoaUsecase {
	return PessoaUsecase{
		pessoaRepository:      pessoarepo,
		abreviaturaRepository: abreviaturarepo,
	}
}

func (pu *PessoaUsecase) GetPessoas() ([]*model.Pessoa, error) {

	pessoas, err := pu.pessoaRepository.GetPessoas()

	if err != nil {
		return nil, err
	}

	for _, values := range pessoas {
		values.Abreviaturas, err = pu.abreviaturaRepository.GetAbreviaturasById(values.IdPessoa)
	}

	return pessoas, nil
}

func (pu *PessoaUsecase) CreatePessoa(pessoa *model.Pessoa) error {
	pessoa, err := pu.pessoaRepository.CreatePessoa(pessoa)

	if err != nil {
		return err
	}

	for _, value := range pessoa.Abreviaturas {

		value.IdPessoa = pessoa.IdPessoa
		err = pu.abreviaturaRepository.CreateAbreviatura(value)

		if err != nil {
			return err
		}
	}

	return nil
}

func (pu *PessoaUsecase) GetPessoaByIdLattes(idLattes int) (*model.Pessoa, error) {

	pessoa, err := pu.pessoaRepository.GetPessoaByIdLattes(idLattes)

	if err != nil {
		return nil, err
	}

	pessoa.Abreviaturas, err = pu.abreviaturaRepository.GetAbreviaturasById(pessoa.IdPessoa)

	if err != nil {
		return nil, err
	}

	return pessoa, nil
}

func (pu *PessoaUsecase) UpdatePessoa(pessoa *model.Pessoa) error {

	err := pu.pessoaRepository.UpdatePessoa(pessoa)

	if err != nil {
		return err
	}

	for _, values := range pessoa.Abreviaturas {
		err = pu.abreviaturaRepository.UpdateAbreviaturas(values)
		
		if err != nil{
			return err
		}
	}
	return nil
}
