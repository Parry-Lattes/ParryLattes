package main

import (
	"net/http"

	"github.com/labstack/echo"

	"parry_end/controllers"
	"parry_end/db"
	"parry_end/repository"
	"parry_end/usecase"
)

func main() {
	e := echo.New()
	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	dbConnection := db.GetDBHandle()

	PessoaRepository := repository.NewPessoaRepository(dbConnection)
	CurriculoRepository := repository.NewCurriculoRepository(dbConnection)
	ProducaoRepository := repository.NewProducaoRepository(dbConnection)
	AbreviaturaRepository := repository.NewAbreviaturaRepository(dbConnection)

	curriculoUsecase := usecase.NewCurriculoUseCase(
		&CurriculoRepository,
		&ProducaoRepository,
		&AbreviaturaRepository,
	)

	pessoaUseCase := usecase.NewPessoaUseCase(
		&PessoaRepository,
		&AbreviaturaRepository,
	)

	pessoaCurriculoUsecase := usecase.NewPessoaCurriculoUsecase(
		&pessoaUseCase,
		&curriculoUsecase,
	)

	dashboardUsecase := usecase.NewDashboardUsecase(
		&CurriculoRepository,
		&ProducaoRepository,
	)

	controllerPessoa := controllers.NewControllerPessoa(&pessoaUseCase)
	controllerCurriculo := controllers.NewControllerCurriculo(&curriculoUsecase)
	controllerPessoaCurriculo := controllers.NewControllerPessoaCurriculo(
		&pessoaCurriculoUsecase,
	)
	controllerDashboard := controllers.NewControllerDashboard(&dashboardUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET(
		"/pessoa",
		controllerPessoa.GetPessoas,
	)
	e.GET(
		"/pessoa/:idLattes",
		controllerPessoa.GetPessoaByIdLattes,
	)
	e.POST(
		"/pessoa",
		controllerPessoa.CreatePessoa,
	)
	// e.PUT("/pessoa", controllerPessoa.UpdatePessoa)
	e.DELETE(
		"/pessoa/:idLattes",
		controllerPessoaCurriculo.DeletePessoa,
	)

	e.GET(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.GetCurriculoById,
	)
	e.GET(
		"/curriculo",
		controllerCurriculo.GetCurriculos,
	)
	e.POST(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.CreateCurriculo,
	)
	// e.PUT("/pessoa/:idLattes/curriculo", controllerCurriculo.UpdateCurriculo)
	e.DELETE(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.DeleteCurriculo,
	)

	e.GET(
		"/dashboard", controllerDashboard.GetRelatorioCompleto,
	)
	e.Logger.Fatal(e.Start(":1323"))
}
