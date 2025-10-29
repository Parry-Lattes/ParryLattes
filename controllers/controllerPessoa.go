package controllers

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type ControllerPessoa struct {
	PessoaUsecase *usecase.PessoaUsecase
}

func NewControllerPessoa(usecase_pessoa *usecase.PessoaUsecase) ControllerPessoa {
	return ControllerPessoa{
		PessoaUsecase: usecase_pessoa,
	}
}

func (c *ControllerPessoa) GetPessoas(e echo.Context) error {

	pessoas, err := c.PessoaUsecase.GetPessoas()

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, pessoas)
}

func (c *ControllerPessoa) GetPessoaByCPF(e echo.Context) error {

	CPF := e.Param("CPF")

	if CPF == "" {

		response := model.Response{
			Message: "Null CPF",
		}
		return e.JSON(http.StatusBadRequest, response)

	}

	CPFPessoa, err := strconv.Atoi(CPF)

	if err != nil {
		response := model.Response{
			Message: "CPF Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	pessoa, err := c.PessoaUsecase.GetPessoaByCPF(CPFPessoa)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	if pessoa == nil {
		response := model.Response{
			Message: "Pessoa not found",
		}
		return e.JSON(http.StatusNotFound, response)
	}

	return e.JSON(http.StatusOK, pessoa)
}

func (c *ControllerPessoa) CreatePessoa(e echo.Context) error {

	var pessoa model.Pessoa
	err := e.Bind(&pessoa)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = c.PessoaUsecase.CreatePessoa(&pessoa)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.NoContent(http.StatusCreated)

}

func (c *ControllerPessoa) UpdatePessoa(e echo.Context) error {

	var pessoa model.Pessoa
	err := e.Bind(&pessoa)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = c.PessoaUsecase.UpdatePessoa(&pessoa)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, nil)
}
