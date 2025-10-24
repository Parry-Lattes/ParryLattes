package repository

import (
	"database/sql"
	"fmt"
	"parry_end/model"
)

type PessoaRepository struct {
	Repository
}

func NewPessoaRepository(connection *sql.DB) PessoaRepository {
	return PessoaRepository{
		Repository: Repository{
			Connection: connection},
	}
}

func (pr *PessoaRepository) GetPessoas() (*[]model.Pessoa, error) {

	querry := "SELECT * FROM Pessoa"
	rows, err := pr.Connection.Query(querry)
	if err != nil {
		fmt.Println(err)
		return &[]model.Pessoa{}, err
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
			return &[]model.Pessoa{}, err
		}

		pessoaList = append(pessoaList, pessoaObj)
	}

	rows.Close()

	return &pessoaList, nil
}

func (pr *PessoaRepository) CreatePessoa(pessoa *model.Pessoa) (int, error) {
	query, err := pr.Connection.Prepare(`
		INSERT INTO Pessoa (Nome, Sexo, Abreviatura, Nacionalidade) 
		VALUES (?, ?, ?, ?)`)

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(
		pessoa.Nome,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao obter último ID:", err)
		return 0, err
	}

	return int(id), nil
}

func (pr *PessoaRepository) GetPessoaById(idPessoa int) (*model.Pessoa, error) {
	query, err := pr.Connection.Prepare("SELECT * FROM Pessoa WHERE idPessoa = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var pessoa model.Pessoa

	err = query.QueryRow(idPessoa).Scan(
		&pessoa.IdPessoa,
		&pessoa.Nome,
		&pessoa.Sexo,
		&pessoa.Abreviatura,
		&pessoa.Nacionalidade)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &pessoa, nil
}

func (pr *PessoaRepository) UpdatePessoa(pessoa *model.Pessoa) error {
	query, err := pr.Connection.Prepare("UPDATE Pessoa " +
		"SET Nome = ?, " +
		"Sexo = ?, " +
		"Abreviatura = ?, " +
		"Nacionalidade = ? " +
		"WHERE idPessoa = ?")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	result, err := query.Exec(
		pessoa.Nome,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
		pessoa.IdPessoa,
	)

	fmt.Println(result)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil

}
