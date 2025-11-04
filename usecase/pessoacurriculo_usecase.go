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

	pessoa, err := cu.PessoaUsecase.GetPessoaByIdLattes(pessoaCurriculo.Pessoa.IdLattes)

	if err != nil {
		fmt.Println("Erro 1")
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.CurriculoUsecase.CurriculoRepository.CreateCurriculo(pessoaCurriculo.Curriculo, pessoa)

	if err != nil {
		fmt.Println("Erro 2")
		return err
	}

	for _, value := range *pessoaCurriculo.Curriculo.Producoes {

		Producao, err := cu.CurriculoUsecase.ProducaoRepository.CreateProducao(&value, pessoaCurriculo.Curriculo)

		if err != nil {
			fmt.Println("Erro 3")
			return err
		}
		err = cu.CurriculoUsecase.CurriculoRepository.LinkCurriculoProducao(pessoaCurriculo.Curriculo, Producao)

		if err != nil {
			fmt.Println("Erro 4")
			return err
		}
	}

	return nil
}
