package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"parry_end/usecase"
)

type ControllerDashboard struct {
	DashboardUsecase *usecase.DashboardUsecase
}

func NewControllerDashboard(
	usecase_dashboard *usecase.DashboardUsecase,
) ControllerDashboard {
	return ControllerDashboard{
		DashboardUsecase: usecase_dashboard,
	}
}

func (c *ControllerDashboard) GetRelatorioCompleto(e echo.Context) error {
	relatorioGeral, err := c.DashboardUsecase.GetRelatorioGeral()
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	err = c.DashboardUsecase.ConstructRelatorioAno(relatorioGeral)
	if err != nil {
		fmt.Println(err)
		e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, relatorioGeral)
}
