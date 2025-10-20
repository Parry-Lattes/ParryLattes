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
	pessoaUseCase := usecase.NewPessoaUseCase(PessoaRepository)
	pessoaController := controller.NewPessoaController(pessoaUseCase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Sexo")
	})

	e.GET("/getpessoa", pessoaController.GetPessoas)

	e.Logger.Fatal(e.Start(":1323"))
}
