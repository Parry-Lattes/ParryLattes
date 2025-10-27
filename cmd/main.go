package main

import (
	"net/http"
	"parry_end/controller"
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

	curriculoUsecase := usecase.NewCurriculoUseCase(CurriculoRepository)
	pessoaUseCase := usecase.NewPessoaUseCase(PessoaRepository)
	producaoUsecase := usecase.NewProducaoUseCase(ProducaoRepository)

	controller := controller.NewController(pessoaUseCase, curriculoUsecase, producaoUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET("/pessoa", controller.GetPessoas)
	e.GET("/pessoa/:idPessoa", controller.GetPessoaById)
	e.POST("/pessoa/create", controller.CreatePessoa)
	e.POST("/pessoa/update", controller.UpdatePessoa)

	e.GET("/curriculo/:idCurriculo", controller.GetCurriculoById)
	e.GET("/curriculo", controller.GetCurriculos)
	e.POST("/curriculo/create", controller.CreateCurriculo)
	e.POST("/curriculo/update", controller.UpdateCurriculo)

	e.GET("/producao", controller.GetProducoes)
	e.Logger.Fatal(e.Start(":1323"))
}
