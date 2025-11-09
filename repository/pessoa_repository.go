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

	querry := "SELECT idPessoa,Nome,idLattes,Nacionalidade " +
		"FROM Pessoa"
	rows, err := pr.Connection.Query(querry)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var pessoaList []model.Pessoa
	var pessoaObj model.Pessoa

	for rows.Next() {
		err = rows.Scan(
			&pessoaObj.IdPessoa,
			&pessoaObj.Nome,
			&pessoaObj.IdLattes,
			&pessoaObj.Nacionalidade,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		pessoaObj.Abreviaturas, err = pr.GetAbreviaturasById(&pessoaObj.IdPessoa)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		pessoaList = append(pessoaList, pessoaObj)
	}

	return &pessoaList, nil
}

func (pr *PessoaRepository) CreatePessoa(pessoa *model.Pessoa) error {
	query, err := pr.Connection.Prepare("INSERT INTO Pessoa (Nome,idLattes,Nacionalidade) " +
		"VALUES (?, ?, ?)")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(
		pessoa.Nome,
		pessoa.IdLattes,
		pessoa.Nacionalidade,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	for _, value := range *pessoa.Abreviaturas {

		value.IdPessoa = pessoa.IdPessoa
		err = pr.CreateAbreviatura(&value)

		if err != nil {
			fmt.Println("Erro ao cadastrar abreviatura", err)
			return err
		}

	}

	return nil
}

func (pr *PessoaRepository) GetPessoaByIdLattes(IdLattes int) (*model.Pessoa, error) {

	query, err := pr.Connection.Prepare("SELECT idPessoa,Nome,idLattes,Nacionalidade FROM Pessoa WHERE idLattes = ?")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var pessoa model.Pessoa

	err = query.QueryRow(IdLattes).Scan(
		&pessoa.IdPessoa,
		&pessoa.Nome,
		&pessoa.IdLattes,
		&pessoa.Nacionalidade)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		fmt.Println(err)
		return nil, err
	}

	pessoa.Abreviaturas, err = pr.GetAbreviaturasById(&pessoa.IdPessoa)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pessoa, nil
}

func (pr *PessoaRepository) UpdatePessoa(pessoa *model.Pessoa) error {

	query, err := pr.Connection.Prepare("UPDATE Pessoa " +
		"SET Nome = ?, " +
		"idLattes = ?, " +
		"Nacionalidade = ? " +
		"WHERE idLattes = ?")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(
		pessoa.Nome,
		pessoa.IdLattes,
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

func (pr *PessoaRepository) GetAbreviaturasById(IdPessoa *int64) (*[]model.Abreviatura, error) {

	query := "SELECT idAbreviatura, idPessoa, abreviatura FROM AbreviaturaPessoa WHERE idPessoa = ?"

	rows, err := pr.Connection.Query(query, IdPessoa)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var abreviaturaList []model.Abreviatura
	var abreviaturaObj model.Abreviatura

	for rows.Next() {
		err = rows.Scan(
			&abreviaturaObj.IdAbreviatura,
			&abreviaturaObj.IdPessoa,
			&abreviaturaObj.Abreviatura,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			fmt.Println(err)
			return nil, err
		}

		abreviaturaList = append(abreviaturaList, abreviaturaObj)
	}

	return &abreviaturaList, nil
}

func (pr *PessoaRepository) CreateAbreviatura(abreviatura *model.Abreviatura) error {

	query, err := pr.Connection.Prepare("INSERT INTO Abreviatura (idPessoa,Abreviatura) VALUES (?,?)")

	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return err
	}

	defer query.Close()

	result, err := query.Exec(
		abreviatura.IdPessoa,
		abreviatura.Abreviatura,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Errro:", err)
		return err
	}

	abreviatura.IdAbreviatura = id

	return nil

}
