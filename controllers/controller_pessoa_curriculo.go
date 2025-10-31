package controllers

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"

	"github.com/labstack/echo"
)

type ControllerPessoaCurriculo struct {
	PessoaCurriculoUsecase *usecase.PessoaCurriculoUsecasse
}

func NewControllerPessoaCurriculo(usecase *usecase.PessoaCurriculoUsecasse) ControllerPessoaCurriculo {
	return ControllerPessoaCurriculo{
		PessoaCurriculoUsecase: usecase,
	}
}

func (pc *ControllerPessoaCurriculo) CreateCurriculo(e echo.Context) error {

	var pessoaCurriculo model.PessoaCurriculo

	err := e.Bind(&pessoaCurriculo)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = pc.PessoaCurriculoUsecase.CreateCurriculo(&pessoaCurriculo)


	return e.JSON(http.StatusOK, err)
}
