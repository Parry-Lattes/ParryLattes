package usecase

import (
	"fmt"

	"parry_end/model"
)

type PessoaCurriculoUsecase struct {
	pessoaUsecase    *PessoaUsecase
	curriculoUsecase *CurriculoUsecase
}

func NewPessoaCurriculoUsecase(
	pessoareUsecase *PessoaUsecase,
	curriculoUsecase *CurriculoUsecase,
) PessoaCurriculoUsecase {
	return PessoaCurriculoUsecase{
		pessoaUsecase:    pessoareUsecase,
		curriculoUsecase: curriculoUsecase,
	}
}

func (cu *PessoaCurriculoUsecase) CreateCurriculo(
	pessoaCurriculo *model.PessoaCurriculo,
) error {
	pessoa, err := cu.pessoaUsecase.GetPessoaByIdLattes(
		pessoaCurriculo.Pessoa.IdLattes,
	)
	if err != nil {
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.curriculoRepository.CreateCurriculo(
		pessoaCurriculo.Curriculo,
		pessoa,
	)

	for _, value := range pessoaCurriculo.Curriculo.Producoes {

		value.TipoId = cu.curriculoUsecase.identifyTipoProducao(value)

		Producao, err := cu.curriculoUsecase.producaoRepository.CreateProducao(
			value,
			pessoaCurriculo.Curriculo,
		)
		if err != nil {
			return err
		}
		err = cu.curriculoUsecase.curriculoRepository.LinkCurriculoProducao(
			pessoaCurriculo.Curriculo,
			Producao,
		)
		if err != nil {
			return err
		}

		for _, coautor := range value.Coautores {

			coautor.Abreviatura.IdPessoa = nil
			coautor.IdProducao = value.IdProducao

			fmt.Println(coautor.IdProducao)
			coautor.Abreviatura, err = cu.pessoaUsecase.abreviaturaRepository.CreateAbreviatura(
				coautor.Abreviatura,
			)
			if err != nil {
				return err
			}
			coautor, err = cu.pessoaUsecase.abreviaturaRepository.CreateACoautor(
				coautor,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cu *PessoaCurriculoUsecase) GetCurriculoByIdLattes(
	idLattes string,
) (*model.Curriculo, error) {
	pessoa, err := cu.pessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		return nil, err
	}

	curriculo, err := cu.curriculoUsecase.GetCurriculoById(pessoa.IdPessoa)
	if err != nil {
		return nil, err
	}

	curriculo.Producoes, err = cu.curriculoUsecase.producaoRepository.GetProducaoByIdLattes(
		curriculo,
	)
	if err != nil {
		return nil, err
	}

	for _, value := range curriculo.Producoes {
		value.Coautores, err = cu.curriculoUsecase.abreviaturaRepository.GetCoautoresByIdProducao(
			curriculo.IdCurriculo,
		)
		if err != nil {
			return nil, err
		}

		for _, coautor := range value.Coautores {
			coautor.Abreviatura, err = cu.curriculoUsecase.abreviaturaRepository.GetAbreviaturaByCoautor(
				coautor,
			)
			if err != nil {
				return nil, err
			}

		}
	}

	return curriculo, nil
}

func (cu *PessoaCurriculoUsecase) DeleteCurriculo(idLattes string) error {
	var pessoaCurriculo *model.PessoaCurriculo = &model.PessoaCurriculo{}
	var err error

	pessoaCurriculo.Pessoa, err = cu.pessoaUsecase.GetPessoaByIdLattes(
		idLattes,
	)
	if err != nil {
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.GetCurriculoById(
		pessoaCurriculo.Pessoa.IdPessoa,
	)
	if err != nil {
		return err
	}

	err = cu.curriculoUsecase.DeleteCoautoresByIdProducao(
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	if err != nil {
		return err
	}

	err = cu.curriculoUsecase.DeleteProducaoByIdCurriculo(
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	if err != nil {
		return err
	}

	err = cu.curriculoUsecase.DeleteCurriculoByIdPessoa(
		pessoaCurriculo.Pessoa.IdPessoa,
	)
	if err != nil {
		return nil
	}

	return nil
}

func (cu *PessoaCurriculoUsecase) DeletePessoa(idLattes string) error {
	var pessoaCurriculo *model.PessoaCurriculo = &model.PessoaCurriculo{}
	var err error

	pessoaCurriculo.Pessoa, err = cu.pessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		return err
	}

	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.GetCurriculoById(
		pessoaCurriculo.Pessoa.IdPessoa,
	)

	if err != nil {
		if pessoaCurriculo.Curriculo == nil {

			erro := cu.pessoaUsecase.DeletePessoa(
				pessoaCurriculo.Pessoa.IdLattes,
			)

			if erro != nil {
				return err
			}

			return nil
		}
	} else {
		erro := cu.DeleteCurriculo(pessoaCurriculo.Pessoa.IdLattes)

		if erro != nil {
			return erro
		}

		erro = cu.pessoaUsecase.DeletePessoa(pessoaCurriculo.Pessoa.IdLattes)
		if erro != nil {
			return erro
		}

		return nil
	}

	return nil
}

// func (cu *PessoaCurriculoUsecase) DeletePessoa(pessoaCurriculo *model.PessoaCurriculo) error {
//
// 	var err error
//
// 	pessoaCurriculo.Pessoa, err = cu.pessoaUsecase.GetPessoaByIdLattes(pessoaCurriculo.Pessoa.IdLattes)
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
