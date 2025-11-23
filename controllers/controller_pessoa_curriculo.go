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

func (c *ControllerPessoaCurriculo) GetCurriculoById(e echo.Context) error {

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

	curriculo, err := c.PessoaCurriculoUsecase.GetCurriculoByIdLattes(idLattes)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	if curriculo == nil {
		response := model.Response{
			Message: "Curriculo not found",
		}
		return e.JSON(http.StatusNotFound, response)
	}

	return e.JSON(http.StatusOK, curriculo)
}


func (pc *ControllerPessoaCurriculo) CreateCurriculo(e echo.Context) error {

	var pessoaCurriculo model.PessoaCurriculo = model.PessoaCurriculo{}

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
	

	err = e.Bind(&pessoaCurriculo.Curriculo)
	fmt.Println(pessoaCurriculo)
	pessoaCurriculo.Pessoa = &model.Pessoa{}
	pessoaCurriculo.Pessoa.IdLattes = idLattes
	fmt.Println("Sexo2")

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	fmt.Println("Sexo3")

	err = pc.PessoaCurriculoUsecase.CreateCurriculo(&pessoaCurriculo)


	return e.JSON(http.StatusOK, err)
}
