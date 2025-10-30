package usecase

import (
	"parry_end/repository"
)

type PessoaCurriculoUsecasse struct {
	PessoaRepository    *repository.PessoaRepository
	CurriculoRepository *repository.CurriculoRepository
}

func NewPessoaCurriculoUsecase(
	pessoarepository *repository.PessoaRepository,
	curriculorepository *repository.CurriculoRepository) PessoaCurriculoUsecasse {
	return PessoaCurriculoUsecasse{
		PessoaRepository:    pessoarepository,
		CurriculoRepository: curriculorepository,
	}
}

// func (pcu *PessoaCurriculoUsecasse) CreateCurriculo(curriculo *model.Curriculo, pessoa *model.Pessoa) error{
// 	idPessoa, err := pcu.GetPesoa

// 	idCurriculo, err := pcu.CurriculoRepository.CreateCurriculo(curriculo,)
// }
