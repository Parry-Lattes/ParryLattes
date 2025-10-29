package controllers

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type ControllerProducao struct {
	ProducaoUsecase *usecase.ProducaoUsecase
}

func NewControllerProducao(usecase_producao *usecase.ProducaoUsecase) ControllerProducao {
	return ControllerProducao{
		ProducaoUsecase: usecase_producao,
	}
}

func (c *ControllerProducao) GetProducoes(e echo.Context) error {

	producao, err := c.ProducaoUsecase.GetProducoes()
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, producao)
}

func (c *ControllerProducao) GetProducaoById(e echo.Context) error {

	id := e.Param("idProducao")

	if id == "" {

		response := model.Response{
			Message: "Null ID",
		}
		return e.JSON(http.StatusBadRequest, response)

	}

	idProducao, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "ID Must be a number",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	producao, err := c.ProducaoUsecase.GetProducaoById(idProducao)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	if producao == nil {
		response := model.Response{
			Message: "Pessoa not found",
		}
		return e.JSON(http.StatusNotFound, response)
	}

	return e.JSON(http.StatusOK, producao)
}
