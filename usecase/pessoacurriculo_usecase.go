package usecase

import (
	"fmt"
	"parry_end/model"
)

type PessoaCurriculoUsecasse struct {
	PessoaUsecase    *PessoaUsecase
	CurriculoUsecase *CurriculoUsecase
}

func NewPessoaCurriculoUsecase(
	pessoareUsecase *PessoaUsecase,
	curriculoUsecase *CurriculoUsecase) PessoaCurriculoUsecasse {
	return PessoaCurriculoUsecasse{
		PessoaUsecase:    pessoareUsecase,
		CurriculoUsecase: curriculoUsecase,
	}
}

func (cu *PessoaCurriculoUsecasse) CreateCurriculo(pessoaCurriculo *model.PessoaCurriculo) error {

	pessoa, err := cu.PessoaUsecase.GetPessoaByCPF(pessoaCurriculo.Pessoa.CPF)

	if err != nil {
		return err
	}

	curriculo, err := cu.CurriculoUsecase.CurriculoRepository.CreateCurriculo(pessoaCurriculo.Curriculo, pessoa)

	if err != nil {
		fmt.Println("Sexo 1")
		return err
	}

	fmt.Println("Sexo")

	for _, value := range *curriculo.Producoes {

		fmt.Println(value)
		Producao, err := cu.CurriculoUsecase.ProducaoRepository.CreateProducao(&value, curriculo)

		if err != nil {
			fmt.Println("Sexo 3")
			return err
		}
		err = cu.CurriculoUsecase.CurriculoRepository.LinkCurriculoProducao(curriculo, Producao)

		if err != nil {
			fmt.Println("Sexo 2")
			return err
		}
	}

	return nil
}
