package controller

import (
	"fmt"
	"net/http"
	"parry_end/usecase"

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
