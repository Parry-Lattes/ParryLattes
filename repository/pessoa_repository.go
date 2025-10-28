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

	querry := "SELECT Nome,CPF,Sexo,Abreviatura,Nacionalidade " +
		"FROM Pessoa"
	rows, err := pr.Connection.Query(querry)
	if err != nil {
		fmt.Println(err)
		return &[]model.Pessoa{}, err
	}

	var pessoaList []model.Pessoa
	var pessoaObj model.Pessoa

	for rows.Next() {
		err = rows.Scan(
			&pessoaObj.Nome,
			&pessoaObj.CPF,
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

func (pr *PessoaRepository) CreatePessoa(pessoa *model.Pessoa) error {
	query, err := pr.Connection.Prepare(`
		INSERT INTO Pessoa (Nome,CPF,Sexo,Abreviatura,Nacionalidade) 
		VALUES (?, ?, ?, ?, ?)`)

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(
		pessoa.Nome,
		pessoa.CPF,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	raffected, err := result.RowsAffected() // Necessário dar uso para result

	if err != nil {
		fmt.Println(raffected)
		fmt.Println("Errro:", err)
		return err
	}

	return nil
}

func (pr *PessoaRepository) GetPessoaByCPF(CPF int) (*model.Pessoa, error) {
	
	query, err := pr.Connection.Prepare("SELECT Nome,CPF,Sexo,Abreviatura,Nacionalidade FROM Pessoa WHERE CPF = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var pessoa model.Pessoa

	err = query.QueryRow(CPF).Scan(
		&pessoa.Nome,
		&pessoa.CPF,
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
		"CPF = ?, " +
		"Sexo = ?, " +
		"Abreviatura = ?, " +
		"Nacionalidade = ? " +
		"WHERE CPF = ?")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	result, err := query.Exec(
		pessoa.Nome,
		pessoa.CPF,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
		pessoa.CPF,
	)

	fmt.Println(result)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil

}
