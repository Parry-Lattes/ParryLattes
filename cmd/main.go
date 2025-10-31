package main

import (
	"net/http"
	"parry_end/controllers"
	"parry_end/db"
	"parry_end/repository"
	"parry_end/usecase"

	"github.com/labstack/echo"
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

	curriculoUsecase := usecase.NewCurriculoUseCase(&CurriculoRepository, &ProducaoRepository)
	pessoaUseCase := usecase.NewPessoaUseCase(&PessoaRepository)
	pessoaCurriculoUsecase := usecase.NewPessoaCurriculoUsecase(&pessoaUseCase, &curriculoUsecase)

	controllerPessoa := controllers.NewControllerPessoa(&pessoaUseCase)
	controllerCurriculo := controllers.NewControllerCurriculo(&curriculoUsecase)
	controllerPessoaCurriculo := controllers.NewControllerPessoaCurriculo(&pessoaCurriculoUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET("/pessoa", controllerPessoa.GetPessoas)
	e.GET("/pessoa/:CPF", controllerPessoa.GetPessoaByCPF)
	e.POST("/pessoa/create", controllerPessoa.CreatePessoa)
	e.POST("/pessoa/update", controllerPessoa.UpdatePessoa)

	e.GET("/curriculo/:idLattes", controllerCurriculo.GetCurriculoById)
	e.GET("/curriculo", controllerCurriculo.GetCurriculos)
	e.POST("/curriculo/create", controllerPessoaCurriculo.CreateCurriculo)
	e.POST("/curriculo/update", controllerCurriculo.UpdateCurriculo)

	e.Logger.Fatal(e.Start(":1323"))
}
