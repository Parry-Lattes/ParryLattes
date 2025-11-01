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

	fmt.Println("sexo")

	// fmt.Println(pessoaCurriculo.Pessoa.CPF)

	pessoa, err := cu.PessoaUsecase.GetPessoaByCPF(pessoaCurriculo.Pessoa.CPF)

	if err != nil {
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.CurriculoUsecase.CurriculoRepository.CreateCurriculo(pessoaCurriculo.Curriculo, pessoa)

	if err != nil {
		return err
	}

	for _, value := range *pessoaCurriculo.Curriculo.Producoes {

		Producao, err := cu.CurriculoUsecase.ProducaoRepository.CreateProducao(&value, pessoaCurriculo.Curriculo)

		if err != nil {
			return err
		}
		err = cu.CurriculoUsecase.CurriculoRepository.LinkCurriculoProducao(pessoaCurriculo.Curriculo, Producao)

		if err != nil {
			return err
		}
	}

	return nil
}
