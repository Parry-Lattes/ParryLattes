package controller

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type Controller struct {
	PessoaUsecase    usecase.PessoaUsecase
	CurriculoUsecase usecase.CurriculoUsecase
	ProducaoUsecase  usecase.ProducaoUsecase
}

func NewController(usecase_pessoa usecase.PessoaUsecase, usecase_curriculo usecase.CurriculoUsecase, usecase_producao usecase.ProducaoUsecase) Controller {
	return Controller{
		PessoaUsecase:    usecase_pessoa,
		CurriculoUsecase: usecase_curriculo,
		ProducaoUsecase:  usecase_producao,
	}
}

func (p *Controller) GetCurriculos(c echo.Context) error {
	curriculos, err := p.CurriculoUsecase.GetCurriculos()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, curriculos)
}

func (p *Controller) GetCurriculoById(c echo.Context) error {

	id := c.Param("idCurriculo")

	if id == "" {
		response := model.Response{
			Message: "Null ID",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	idCurrculo, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	curriculo, err := p.CurriculoUsecase.GetCurriculoById(idCurrculo)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	if curriculo == nil {
		response := model.Response{
			Message: "Curriculo not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, curriculo)
}

func (p *Controller) GetPessoas(c echo.Context) error {

	pessoas, err := p.PessoaUsecase.GetPessoas()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, pessoas)
}

func (p *Controller) CreateCurriculo(c echo.Context) error {

	var curriculo model.Curriculo
	err := c.Bind(&curriculo)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	var id int
	insertedCurriculo, err := p.CurriculoUsecase.CreateCurriculo(&curriculo, id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, insertedCurriculo)

}

func (p *Controller) CreatePessoa(c echo.Context) error {

	var pessoa model.Pessoa
	err := c.Bind(&pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	err = p.PessoaUsecase.CreatePessoa(&pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)

}

func (p *Controller) GetPessoaByCPF(c echo.Context) error {

	CPF := c.Param("CPF")

	if CPF == "" {

		response := model.Response{
			Message: "Null CPF",
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	CPFPessoa, err := strconv.Atoi(CPF)

	if err != nil {
		response := model.Response{
			Message: "CPF Must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	pessoa, err := p.PessoaUsecase.GetPessoaByCPF(CPFPessoa)

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

func (p *Controller) UpdatePessoa(c echo.Context) error {

	var pessoa model.Pessoa
	err := c.Bind(&pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	err = p.PessoaUsecase.UpdatePessoa(&pessoa)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *Controller) UpdateCurriculo(c echo.Context) error {
	var curriculo model.Curriculo
	err := c.Bind(&curriculo)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
	}

	err = p.CurriculoUsecase.UpdateCurriculo(&curriculo)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)

}

func (p *Controller) GetProducoes(c echo.Context) error {

	producao, err := p.ProducaoUsecase.GetProducoes()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, producao)
}

func (p *Controller) GetProducaoById(c echo.Context) error {

	id := c.Param("idProducao")

	if id == "" {

		response := model.Response{
			Message: "Null ID",
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	idProducao, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	producao, err := p.ProducaoUsecase.GetProducaoById(idProducao)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	}

	if producao == nil {
		response := model.Response{
			Message: "Pessoa not found",
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, producao)
}
