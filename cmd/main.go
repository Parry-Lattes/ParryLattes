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
	curriculoUsecase := usecase.NewCurriculoUseCase(CurriculoRepository)
	pessoaUseCase := usecase.NewPessoaUseCase(PessoaRepository)
	controller := controller.NewController(pessoaUseCase, curriculoUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET("/curriculo", controller.GetCurriculos)
	e.GET("/pessoa", controller.GetPessoas)
	e.GET("/pessoa/:idPessoa", controller.GetPessoaById)
	e.GET("/curriculo/:idCurriculo", controller.GetCurriculoById)
	e.POST("/createpessoa", controller.CreatePessoa)

	e.Logger.Fatal(e.Start(":1323"))
}
