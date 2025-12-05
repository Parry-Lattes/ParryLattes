package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"parry_end/controllers"
	"parry_end/db"
	"parry_end/model"
	"parry_end/repository"
	"parry_end/usecase"
)

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("Error generating token")
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

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



	auth := e.Group("/v1")

	auth.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET(
	routes.GET(
		"/pessoa",
		controllerPessoa.GetPessoas,
	)
	routes.GET(
		"/pessoa/:idLattes",
		controllerPessoa.GetPessoaByIdLattes,
	)
	routes.POST(
		"/pessoa",
		controllerPessoa.CreatePessoa,
	)
	// routes.PUT("/pessoa", controllerPessoa.UpdatePessoa)
	routes.DELETE(
		"/pessoa/:idLattes",
		controllerPessoaCurriculo.DeletePessoa,
	)

	routes.GET(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.GetCurriculoById,
	)
	routes.POST(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.CreateCurriculo,
	)
	// routes.PUT("/pessoa/:idLattes/curriculo", controllerCurriculo.UpdateCurriculo)
	routes.DELETE(
		"/pessoa/:idLattes/curriculo",
		controllerPessoaCurriculo.DeleteCurriculo,
	)

	routes.GET(
		"/curriculo",
		controllerCurriculo.GetCurriculos,
	)

	routes.GET("/dashboard", controllerDashboard.GetRelatorioCompleto)

	e.Logger.Fatal(e.Start(":1323"))
}
