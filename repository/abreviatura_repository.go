package repository

import (
	"database/sql"
	"fmt"

	"parry_end/model"
)

type AbreviaturaRepository struct {
	Repository
}

func NewAbreviaturaRepository(connection *sql.DB) AbreviaturaRepository {
	return AbreviaturaRepository{
		Repository: Repository{
			Connection: connection,
		},
	}
}

func (ar *AbreviaturaRepository) GetAbreviaturasById(
	IdPessoa int64,
) ([]*model.Abreviatura, error) {
	query := "SELECT idAbreviatura, idPessoa, abreviatura FROM Abreviatura WHERE idPessoa = ?"

	rows, err := ar.Connection.Query(query, IdPessoa)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var abreviaturaList []*model.Abreviatura

	for rows.Next() {

		var abreviaturaObj model.Abreviatura = model.Abreviatura{}

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

		abreviaturaList = append(abreviaturaList, &abreviaturaObj)
	}

	return abreviaturaList, nil
}

func (ar *AbreviaturaRepository) GetAbreviaturaByCoautor(
	coautor *model.Coautor,
) (*model.Abreviatura, error) {
	query, err := ar.Connection.Prepare(
		"SELECT idAbreviatura,idPessoa,Abreviatura " +
			"FROM Abreviatura " +
			"WHERE idAbreviatura = ?",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var abreviatura *model.Abreviatura = &model.Abreviatura{}

	err = query.QueryRow(coautor.Abreviatura.IdAbreviatura).Scan(
		&abreviatura.IdAbreviatura,
		&abreviatura.IdPessoa,
		&abreviatura.Abreviatura,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	return abreviatura, nil
}

func (ar *AbreviaturaRepository) GetCoautoresByIdProducao(
	idProducao int64,
) ([]*model.Coautor, error) {
	query := "SELECT idCoautor,idProducao,idAbreviatura " +
		"FROM Coautor " +
		"WHERE idProducao = ?"
	rows, err := ar.Connection.Query(query, idProducao)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var CoautorList []*model.Coautor

	for rows.Next() {

		var CoautorObj model.Coautor = model.Coautor{Abreviatura: &model.Abreviatura{}}

		err = rows.Scan(
			&CoautorObj.IdCoautor,
			&CoautorObj.IdProducao,
			&CoautorObj.Abreviatura.IdAbreviatura,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			fmt.Println(err)
			return nil, err
		}

		CoautorList = append(CoautorList, &CoautorObj)
	}

	return CoautorList, nil
}

func (ar *AbreviaturaRepository) CreateAbreviatura(
	abreviatura *model.Abreviatura,
) (*model.Abreviatura, error) {
	query, err := ar.Connection.Prepare(
		"INSERT INTO Abreviatura (idPessoa,Abreviatura) VALUES (?,?)",
	)
	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return nil, err
	}

	defer query.Close()

	result, err := query.Exec(
		abreviatura.IdPessoa,
		abreviatura.Abreviatura,
	)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Errro:", err)
		return nil, err
	}

	abreviatura.IdAbreviatura = id

	return abreviatura, nil
}

func (ar *AbreviaturaRepository) UpdateAbreviaturas(
	abreviatura *model.Abreviatura,
) error {
	query, err := ar.Connection.Prepare("UPDATE Abreviatura " +
		"SET idPessoa = ? " +
		"SET Abreviatura = ? " +
		"HWERE idAbreviatura = ?")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(
		abreviatura.IdPessoa,
		abreviatura.Abreviatura,
		abreviatura.IdAbreviatura,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (ar AbreviaturaRepository) CreateACoautor(
	coautor *model.Coautor,
) (*model.Coautor, error) {
	query, err := ar.Connection.Prepare(
		"INSERT INTO Coautor (idProducao,idAbreviatura) VALUES (?,?)",
	)
	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return nil, err
	}

	defer query.Close()

	result, err := query.Exec(
		coautor.IdProducao,
		coautor.Abreviatura.IdAbreviatura,
	)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao obter o último ID:", err)
		return nil, err
	}

	coautor.IdCoautor = id

	return coautor, nil
}
