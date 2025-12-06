package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"parry_end/controllers"
	"parry_end/db"
	"parry_end/middlewares"
	"parry_end/repository"
	"parry_end/usecase"
)

func main() {
	engine := echo.New()
	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	dbConnection := db.GetDBHandle()

	SessaoRepository := repository.NewSessaoRepository(dbConnection)
	LoginRepository := repository.NewLoginRepository(dbConnection)
	PessoaRepository := repository.NewPessoaRepository(dbConnection)
	CurriculoRepository := repository.NewCurriculoRepository(dbConnection)
	ProducaoRepository := repository.NewProducaoRepository(dbConnection)
	AbreviaturaRepository := repository.NewAbreviaturaRepository(dbConnection)

	loginUsecase := usecase.NewLoginUsecase(
		&LoginRepository,
		&SessaoRepository,
	)

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

	authMiddleware := middlewares.NewAuthMiddleware(&loginUsecase)

	controllerLogin := controllers.NewControllerLogin(&loginUsecase)
	controllerPessoa := controllers.NewControllerPessoa(&pessoaUseCase)
	controllerCurriculo := controllers.NewControllerCurriculo(&curriculoUsecase)
	controllerPessoaCurriculo := controllers.NewControllerPessoaCurriculo(
		&pessoaCurriculoUsecase,
	)
	controllerDashboard := controllers.NewControllerDashboard(&dashboardUsecase)

	routes := engine.Group("/v1")

	routes.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})
	// auth.Use()

	routes.POST("/login", controllerLogin.LoginUser)

	protected := routes.Group("")

	protected.Use(authMiddleware.CheckIfCSRFTokenExists)
	protected.Use(authMiddleware.CheckIfSessionIsValid)

	protected.GET(
		"/pessoas",
		controllerPessoa.GetPessoas,
	)
	protected.GET(
		"/pessoas/:idLattes",
		controllerPessoa.GetPessoaByIdLattes,
	)
	protected.POST(
		"/pessoas",
		controllerPessoa.CreatePessoa,
	)
	// routes.PUT("/pessoas", controllerPessoa.UpdatePessoa)
	protected.DELETE(
		"/pessoas/:idLattes",
		controllerPessoaCurriculo.DeletePessoa,
	)

	protected.GET(
		"/pessoas/:idLattes/curriculo",
		controllerPessoaCurriculo.GetCurriculoById,
	)
	protected.POST(
		"/pessoas/:idLattes/curriculo",
		controllerPessoaCurriculo.CreateCurriculo,
	)
	// routes.PUT("/pessoas/:idLattes/curriculo", controllerCurriculo.UpdateCurriculo)
	protected.DELETE(
		"/pessoas/:idLattes/curriculo",
		controllerPessoaCurriculo.DeleteCurriculo,
	)

	protected.GET(
		"/curriculos",
		controllerCurriculo.GetCurriculos,
	)

	protected.GET("/dashboard", controllerDashboard.GetRelatorioCompleto)

	engine.Logger.Fatal(engine.Start(":1323"))
}
