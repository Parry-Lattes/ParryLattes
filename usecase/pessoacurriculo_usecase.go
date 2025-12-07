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
	fmt.Println("Pegando Pessoa:", pessoaCurriculo.Pessoa)
	pessoa, err := cu.pessoaUsecase.GetPessoaByIdLattes(
		pessoaCurriculo.Pessoa.IdLattes,
	)
	if err != nil {
		return err
	}

	fmt.Println("Pegando Curriculo da Pessoa:", pessoa)
	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.curriculoRepository.CreateCurriculo(
		pessoaCurriculo.Curriculo,
		pessoa,
	)

	for _, value := range pessoaCurriculo.Curriculo.Producoes {

		fmt.Println("Identificando Tipo da Producao:", value)
		value.TipoId = cu.curriculoUsecase.identifyTipoProducao(value)

		fmt.Println("Criando Producao:", value)
		Producao, err := cu.curriculoUsecase.producaoRepository.CreateProducao(
			value,
			pessoaCurriculo.Curriculo,
		)
		if err != nil {
			return err
		}

		fmt.Println(
			"Linkando Curriculo",
			pessoaCurriculo.Curriculo,
			"à Producao:",
			value,
		)
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

			fmt.Println("Registrando Criando Abreviatura:", coautor.Abreviatura)
			coautor.Abreviatura, err = cu.pessoaUsecase.abreviaturaRepository.CreateAbreviatura(
				coautor.Abreviatura,
			)
			if err != nil {
				return err
			}

			fmt.Println("Criando Coautor Para a Abreviatura", coautor)
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
	fmt.Println("Pegando pessoa por id Lattes", idLattes)
	pessoa, err := cu.pessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		return nil, err
	}

	fmt.Println("Pegando Curriculo da Pessoa", pessoa)
	curriculo, err := cu.curriculoUsecase.GetCurriculoById(pessoa.IdPessoa)
	if err != nil {
		return nil, err
	}

	fmt.Println("Pegando Producoes do curriculo", curriculo)
	curriculo.Producoes, err = cu.curriculoUsecase.producaoRepository.GetProducaoByIdLattes(
		curriculo,
	)
	if err != nil {
		return nil, err
	}

	for _, value := range curriculo.Producoes {

		fmt.Println("Pegando Coautores por Id Producao", value)
		value.Coautores, err = cu.curriculoUsecase.abreviaturaRepository.GetCoautoresByIdProducao(
			value.IdProducao,
		)
		if err != nil {
			return nil, err
		}

		for _, coautor := range value.Coautores {

			fmt.Println("Pegando Abreviatura Por Coautor:", coautor)
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

	fmt.Println("Pegando pessoa por id lattes:", idLattes)
	pessoaCurriculo.Pessoa, err = cu.pessoaUsecase.GetPessoaByIdLattes(
		idLattes,
	)
	if err != nil {
		return err
	}

	fmt.Println(
		"Pegando Curriculo por ID Pessoa",
		pessoaCurriculo.Pessoa.IdPessoa,
	)
	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.GetCurriculoById(
		pessoaCurriculo.Pessoa.IdPessoa,
	)
	if err != nil {
		return err
	}

	fmt.Println(
		"Deletando Coautor Por id Producao:",
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	err = cu.curriculoUsecase.DeleteCoautoresByIdProducao(
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	if err != nil {
		return err
	}

	fmt.Println(
		"Deletando Produao Por Id Curriculo:",
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	err = cu.curriculoUsecase.DeleteProducaoByIdCurriculo(
		pessoaCurriculo.Curriculo.IdCurriculo,
	)
	if err != nil {
		return err
	}

	fmt.Println(
		"Delete Curriculo Por Id Pessoa",
		pessoaCurriculo.Pessoa.IdPessoa,
	)
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

	fmt.Println("Get Pessoa Por Id Lattes:", idLattes)
	pessoaCurriculo.Pessoa, err = cu.pessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		return err
	}

	fmt.Println("Pegando Curriculo Por id:", pessoaCurriculo.Pessoa.IdPessoa)
	pessoaCurriculo.Curriculo, err = cu.curriculoUsecase.GetCurriculoById(
		pessoaCurriculo.Pessoa.IdPessoa,
	)

	if err != nil {
		if pessoaCurriculo.Curriculo == nil {

			fmt.Println("Deletando Pessoa", pessoaCurriculo.Pessoa.IdLattes)
			erro := cu.pessoaUsecase.DeletePessoa(
				pessoaCurriculo.Pessoa.IdLattes,
			)

			if erro != nil {
				return err
			}

			return nil
		}
	} else {

		fmt.Println("Deletando Curriculo", pessoaCurriculo.Pessoa.IdLattes)
		erro := cu.DeleteCurriculo(pessoaCurriculo.Pessoa.IdLattes)

		if erro != nil {
			return erro
		}

		fmt.Println("Deletando Pessoa:", pessoaCurriculo.Pessoa.IdLattes)
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
