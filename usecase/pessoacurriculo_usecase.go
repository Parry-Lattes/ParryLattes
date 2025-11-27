package usecase

import (
	"parry_end/model"
)

type PessoaCurriculoUsecase struct {
	PessoaUsecase    *PessoaUsecase
	CurriculoUsecase *CurriculoUsecase
}

func NewPessoaCurriculoUsecase(
	pessoareUsecase *PessoaUsecase,
	curriculoUsecase *CurriculoUsecase,
) PessoaCurriculoUsecase {
	return PessoaCurriculoUsecase{
		PessoaUsecase:    pessoareUsecase,
		CurriculoUsecase: curriculoUsecase,
	}
}

func (cu *PessoaCurriculoUsecase) CreateCurriculo(
	pessoaCurriculo *model.PessoaCurriculo,
) error {
	pessoa, err := cu.PessoaUsecase.GetPessoaByIdLattes(
		pessoaCurriculo.Pessoa.IdLattes,
	)
	if err != nil {
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.CurriculoUsecase.CurriculoRepository.CreateCurriculo(
		pessoaCurriculo.Curriculo,
		pessoa,
	)
	if err != nil {
		return err
	}

	for _, value := range pessoaCurriculo.Curriculo.Producoes {

		Producao, err := cu.CurriculoUsecase.ProducaoRepository.CreateProducao(
			value,
			pessoaCurriculo.Curriculo,
		)
		if err != nil {
			return err
		}
		err = cu.CurriculoUsecase.CurriculoRepository.LinkCurriculoProducao(
			pessoaCurriculo.Curriculo,
			Producao,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cu *PessoaCurriculoUsecase) GetCurriculoByIdLattes(
	idLattes int,
) (*model.Curriculo, error) {
	pessoa, err := cu.PessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		return nil, err
	}

	curriculo, err := cu.CurriculoUsecase.GetCurriculoById(pessoa.IdPessoa)
	if err != nil {
		return nil, err
	}

	curriculo.Producoes, err = cu.CurriculoUsecase.ProducaoRepository.GetProducaoByIdLattes(
		curriculo,
	)
	if err != nil {
		return nil, err
	}

	for _, value := range curriculo.Producoes {
		value.Coautores, err = cu.CurriculoUsecase.AbreviaturaRepository.GetAbreviaturaByIdProducao(
			value.IdProducao,
		)
	}

	return curriculo, nil
}

// func (cu *PessoaCurriculoUsecase) DeletePessoa(pessoaCurriculo *model.PessoaCurriculo) error {
//
// 	var err error
//
// 	pessoaCurriculo.Pessoa, err = cu.PessoaUsecase.GetPessoaByIdLattes(pessoaCurriculo.Pessoa.IdLattes)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	pessoaCurriculo.Curriculo, err = cu.CurriculoUsecase.GetCurriculoById(pessoaCurriculo.Curriculo.IdPessoa)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	for _, value := range pessoaCurriculo.Curriculo.Producoes {
// 		err = cu.CurriculoUsecase.CurriculoRepository.UnlinkProducaoCurriculo(pessoaCurriculo.Curriculo, value)
// 		if err != nil {
// 			return err
// 		}
//
// 		err = cu.CurriculoUsecase.DeleteProducao(*value)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
//
// }
