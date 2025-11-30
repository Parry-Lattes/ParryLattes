package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"parry_end/model"
	"parry_end/usecase"
)

type ControllerPessoaCurriculo struct {
	PessoaCurriculoUsecase *usecase.PessoaCurriculoUsecase
}

type TipoDeProducao int

const (
	Bibliografica TipoDeProducao = iota
	Patente
	Tecnica
	Outro
)

func NewControllerPessoaCurriculo(
	usecase *usecase.PessoaCurriculoUsecase,
) ControllerPessoaCurriculo {
	return ControllerPessoaCurriculo{
		PessoaCurriculoUsecase: usecase,
	}
}

func (c ControllerPessoaCurriculo) typeConvert(tipo string) TipoDeProducao {
	switch tipo {
	case "Bibliográfica":
		return Bibliografica
	case "Patente":
		return Patente
	case "Técica":
		return Tecnica
	default:
		return Outro
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
		if err == sql.ErrNoRows {
			response := model.Response{
				Message: "Curriculo not found",
			}
			e.JSON(http.StatusNotFound, response)
		}
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
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

	pessoaCurriculo.Pessoa = &model.Pessoa{}
	pessoaCurriculo.Pessoa.IdLattes = idLattes

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = pc.PessoaCurriculoUsecase.CreateCurriculo(&pessoaCurriculo)
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, err)
}
