package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"parry_end/model"
	"parry_end/usecase"
)

type ControllerCurriculo struct {
	curriculoUsecase *usecase.CurriculoUsecase
}

func NewControllerCurriculo(
	usecase_curriculo *usecase.CurriculoUsecase,
) ControllerCurriculo {
	return ControllerCurriculo{
		curriculoUsecase: usecase_curriculo,
	}
}

func (c *ControllerCurriculo) GetCurriculos(e echo.Context) error {
	curriculos, err := c.curriculoUsecase.GetCurriculos()
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, curriculos)
}

func (c *ControllerCurriculo) UpdateCurriculo(e echo.Context) error {
	var curriculo model.Curriculo
	err := e.Bind(&curriculo)
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusBadRequest, err)
	}

	err = c.curriculoUsecase.UpdateCurriculo(&curriculo)
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
// 	insertedCurriculo, err := c.curriculoUsecase.CreateCurriculo(&curriculo, id)

// 	if err != nil {
// 		fmt.Println(err)
// 		e.JSON(http.StatusInternalServerError, err)
// 	}

// 	return e.JSON(http.StatusCreated, insertedCurriculo)

// }
