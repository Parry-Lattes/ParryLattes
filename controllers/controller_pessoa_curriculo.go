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
	pessoaCurriculoUsecase *usecase.PessoaCurriculoUsecase
}

func NewControllerPessoaCurriculo(
	usecase *usecase.PessoaCurriculoUsecase,
) ControllerPessoaCurriculo {
	return ControllerPessoaCurriculo{
		pessoaCurriculoUsecase: usecase,
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

	_, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	curriculo, err := c.pessoaCurriculoUsecase.GetCurriculoByIdLattes(
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			response := model.Response{
				Message: "Curriculo not found",
			}
			return e.JSON(http.StatusNotFound, response)
		}
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
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

	_, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	err = e.Bind(&pessoaCurriculo.Curriculo)

	pessoaCurriculo.Pessoa = &model.Pessoa{}
	pessoaCurriculo.Pessoa.IdLattes = id

	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusBadRequest, err)
	}

	err = pc.pessoaCurriculoUsecase.CreateCurriculo(&pessoaCurriculo)
	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, err)
}

func (pc *ControllerPessoaCurriculo) DeleteCurriculo(e echo.Context) error {
	id := e.Param("idLattes")

	if id == "" {
		response := model.Response{
			Message: "Null ID",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	err = pc.pessoaCurriculoUsecase.DeleteCurriculo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, err)
		}
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, err)
}

func (pc *ControllerPessoaCurriculo) DeletePessoa(e echo.Context) error {
	id := e.Param("idLattes")
	if id == "" {
		response := model.Response{
			Message: "Null ID",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	err = pc.pessoaCurriculoUsecase.DeletePessoa(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return e.JSON(http.StatusNotFound, err)
		}
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, err)
}
