package repository

import (
	"database/sql"
	"fmt"
	"parry_end/model"
)

type PessoaRepository struct {
	connection *sql.DB
}

func NewPessoaRepository(connection *sql.DB) PessoaRepository {
	return PessoaRepository{
		connection: connection,
	}
}

func (pr *PessoaRepository) GetPessoas() ([]model.Pessoa, error) {

	querry := "SELECT * FROM Pessoa"
	rows, err := pr.connection.Query(querry)
	if err != nil {
		fmt.Println(err)
		return []model.Pessoa{}, err
	}

	var pessoaList []model.Pessoa
	var pessoaObj model.Pessoa

	for rows.Next() {
		err = rows.Scan(
			&pessoaObj.IdPessoa,
			&pessoaObj.Nome,
			&pessoaObj.Sexo,
			&pessoaObj.Abreviatura,
			&pessoaObj.Nacionalidade,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Pessoa{}, err
		}

		pessoaList = append(pessoaList, pessoaObj)
	}

	rows.Close()

	return pessoaList, nil
}
