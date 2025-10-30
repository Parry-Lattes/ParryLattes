package controllers

import (
	"fmt"
	"net/http"
	"parry_end/model"
	"parry_end/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type ControllerCurriculo struct {
	CurriculoUsecase *usecase.CurriculoUsecase
}

func NewControllerCurriculo(usecase_curriculo *usecase.CurriculoUsecase) ControllerCurriculo {
	return ControllerCurriculo{
		CurriculoUsecase: usecase_curriculo,
	}
}

func (c *ControllerCurriculo) GetCurriculos(e echo.Context) error {
	curriculos, err := c.CurriculoUsecase.GetCurriculos()
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, curriculos)
}

func (c *ControllerCurriculo) GetCurriculoById(e echo.Context) error {

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

	curriculo, err := c.CurriculoUsecase.GetCurriculoById(idLattes)

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
func (c *ControllerCurriculo) UpdateCurriculo(e echo.Context) error {
	var curriculo model.Curriculo
	err := e.Bind(&curriculo)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = c.CurriculoUsecase.UpdateCurriculo(&curriculo)

	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, nil)

}

// func (c *ControllerCurriculo) CreateCurriculo(e echo.Context) error {

// 	var curriculo model.Curriculo
// 	err := e.Bind(&curriculo)

// 	if err != nil {
// 		fmt.Println(err)
// 		e.JSON(http.StatusBadRequest, err)
// 	}

// 	var id int
// 	insertedCurriculo, err := c.CurriculoUsecase.CreateCurriculo(&curriculo, id)

// 	if err != nil {
// 		fmt.Println(err)
// 		e.JSON(http.StatusInternalServerError, err)
// 	}

// 	return e.JSON(http.StatusCreated, insertedCurriculo)

// }
