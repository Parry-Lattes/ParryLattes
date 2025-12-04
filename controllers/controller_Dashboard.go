package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"parry_end/usecase"
)

type ControllerDashboard struct {
	dashboardUsecase *usecase.DashboardUsecase
}

func NewControllerDashboard(
	usecase_dashboard *usecase.DashboardUsecase,
) ControllerDashboard {
	return ControllerDashboard{
		dashboardUsecase: usecase_dashboard,
	}
}

func (c *ControllerDashboard) GetRelatorioCompleto(e echo.Context) error {
	relatorioGeral, err := c.dashboardUsecase.GetRelatorioGeral()
	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	err = c.dashboardUsecase.ConstructRelatorioAno(relatorioGeral)
	if err != nil {
		fmt.Println(err)
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, relatorioGeral)
}
