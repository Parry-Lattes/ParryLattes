package controller

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type pessoaController struct {
	PessoaUsecase usecase.PessoaUsecase
}

func NewPessoaController(usecase usecase.PessoaUsecase) pessoaController {
	return pessoaController{
		PessoaUsecase: usecase,
	}
}

func (p *pessoaController) GetPessoas(c echo.Context) error {

	pessoas, err := p.PessoaUsecase.GetPessoas()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, pessoas)
}

func (p *pessoaController) CreatePessoa(c echo.Context) error {

	var pessoa model.Pessoa
	err := c.Bind(&pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	insertedPessoa, err := p.PessoaUsecase.CreatePessoa(pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, insertedPessoa)

}

func (p *pessoaController) GetPessoaById(c echo.Context) error {

	id := c.Param("idPessoa")

	if id == "" {

		response := model.Response{
			Message: "Null ID",
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	idPessoa, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	pessoa, err := p.PessoaUsecase.GetPessoaById(idPessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	if pessoa == nil {
		response := model.Response{
			Message: "Pessoa not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, pessoa)
}
