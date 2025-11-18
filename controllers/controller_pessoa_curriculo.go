package controllers

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type ControllerPessoaCurriculo struct {
	PessoaCurriculoUsecase *usecase.PessoaCurriculoUsecase
}

func NewControllerPessoaCurriculo(usecase *usecase.PessoaCurriculoUsecase) ControllerPessoaCurriculo {
	return ControllerPessoaCurriculo{
		PessoaCurriculoUsecase: usecase,
	}
}

func (pc *ControllerPessoaCurriculo) CreateCurriculo(e echo.Context) error {

	var pessoaCurriculo model.PessoaCurriculo

	id := e.Param("idLattes")
	
		if id == "" {
		response := model.Response{
			Message: "Null ID",
		}

		return e.JSON(http.StatusBadRequest, response)
	}

	idLattes, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	pessoaCurriculo.Pessoa.IdLattes = idLattes
	err = e.Bind(&pessoaCurriculo.Curriculo)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = pc.PessoaCurriculoUsecase.CreateCurriculo(&pessoaCurriculo)


	return e.JSON(http.StatusOK, err)
}
