package usecase

import (
	"parry_end/model"
	"parry_end/repository"
)

type CurriculoUsecase struct {
	CurriculoRepository *repository.CurriculoRepository
	ProducaoRepository  *repository.ProducaoRepository
}

func NewCurriculoUseCase(curriculorepo *repository.CurriculoRepository, producaorepo *repository.ProducaoRepository) CurriculoUsecase {
	return CurriculoUsecase{
		CurriculoRepository: curriculorepo,
		ProducaoRepository:  producaorepo,
	}
}

func (cu *CurriculoUsecase) GetCurriculos() (*[]model.Curriculo, error) {
	return cu.CurriculoRepository.GetCurriculos()
}

func (cu *CurriculoUsecase) GetCurriculoById(idLattes int) (*model.Curriculo, error) {

	curriculo, err := cu.CurriculoRepository.GetCurriculoById(idLattes)

	if err != nil {
		return nil, err
	}

	curriculo.Producoes, err = cu.ProducaoRepository.GetProducaoById(curriculo)

	if err != nil {
		return nil, err
	}

	return curriculo, nil
}



func (cu *CurriculoUsecase) UpdateCurriculo(curriculo *model.Curriculo) error {
	
	err := cu.CurriculoRepository.UpdateCurriculo(curriculo)

	if err != nil {
		return err
	}

	return nil
}

