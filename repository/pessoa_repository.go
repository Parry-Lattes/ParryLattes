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

	querry := "SELECT idPessoa,Nome,idLattes,Sexo,Abreviatura,Nacionalidade " +
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
			&pessoaObj.IdPessoa,
			&pessoaObj.Nome,
			&pessoaObj.IdLattes,
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
	query, err := pr.Connection.Prepare("INSERT INTO Pessoa (Nome,idLattes,Sexo,Abreviatura,Nacionalidade) " +
		"VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(
		pessoa.Nome,
		pessoa.IdLattes,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Errro:", err)
		return err
	}

	return nil
}

func (pr *PessoaRepository) GetPessoaByIdLattes(IdLattes int) (*model.Pessoa, error) {

	query, err := pr.Connection.Prepare("SELECT idPessoa,Nome,idLattes,Sexo,Abreviatura,Nacionalidade FROM Pessoa WHERE idLattes = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var pessoa model.Pessoa

	err = query.QueryRow(IdLattes).Scan(
		&pessoa.IdPessoa,
		&pessoa.Nome,
		&pessoa.IdLattes,
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
		"idLattes = ?, " +
		"Sexo = ?, " +
		"Abreviatura = ?, " +
		"Nacionalidade = ? " +
		"WHERE idLattes = ?")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	_, err = query.Exec(
		pessoa.Nome,
		pessoa.IdLattes,
		pessoa.Sexo,
		pessoa.Abreviatura,
		pessoa.Nacionalidade,
		pessoa.IdLattes,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil

}

func (pr *PessoaRepository) DeletePessoa(idLattes int64) error {
	query := "DELETE FROM Pessoa WHERE idLattes = ?"

	result, err := pr.Connection.Exec(query, idLattes)
	if err != nil {
		fmt.Println("erro ao deletar Pessoa")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("erro ao coletar linhas afetadas")
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("pessoa não encontrada:")
		return err
	}

	return nil

}
