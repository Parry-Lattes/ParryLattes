package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"parry_end/model"
	"parry_end/usecase"
)

type ControllerPessoa struct {
	PessoaUsecase *usecase.PessoaUsecase
}

func NewControllerPessoa(
	usecase_pessoa *usecase.PessoaUsecase,
) ControllerPessoa {
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

func (c *ControllerPessoa) GetPessoaByIdLattes(e echo.Context) error {
	idLattes := e.Param("idLattes")

	if idLattes == "" {

		response := model.Response{
			Message: "Null IDLattes",
		}
		return e.JSON(http.StatusBadRequest, response)

	}

	_, err := strconv.Atoi(idLattes)
	if err != nil {
		response := model.Response{
			Message: "IdLattes Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	pessoa, err := c.PessoaUsecase.GetPessoaByIdLattes(idLattes)
	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
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
		fmt.Println("Sedxo",err)
		e.JSON(http.StatusBadRequest, err)
	}

	idLattes := pessoa.IdLattes

	if idLattes == "" {

		response := model.Response{
			Message: "Null IDLattes",
		}
		return e.JSON(http.StatusBadRequest, response)

	}

	_, err = strconv.Atoi(idLattes)
	if err != nil {
		response := model.Response{
			Message: "IdLattes Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	err = c.PessoaUsecase.CreatePessoa(&pessoa)
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.NoContent(http.StatusCreated)
}

// func (c *ControllerPessoa) UpdatePessoa(e echo.Context) error {
// 	var pessoa *model.Pessoa
// 	err := e.Bind(&pessoa)
// 	if err != nil {
// 		fmt.Println(err)
// 		e.JSON(http.StatusBadRequest, err)
// 	}
//
// 	pessoa, err = c.PessoaUsecase.GetPessoaByIdLattes(pessoa.IdLattes)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println(err)
// 			return e.JSON(http.StatusNotFound, err)
// 		}
//
// 		fmt.Println(err)
// 		return e.JSON(http.StatusInternalServerError, err)
// 	}
//
// 	err = c.PessoaUsecase.UpdatePessoa(pessoa)
// 	if err != nil {
// 		fmt.Println(err)
// 		e.JSON(http.StatusInternalServerError, err)
// 	}
//
// 	return e.JSON(http.StatusOK, nil)
// }
