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
			Connection: connection},
	}
}

func (ar *AbreviaturaRepository) GetAbreviaturasById(IdPessoa int64) ([]*model.Abreviatura, error) {

	query := "SELECT idAbreviatura, idPessoa, abreviatura FROM Abreviatura WHERE idPessoa = ?"

	rows, err := ar.Connection.Query(query, IdPessoa)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var abreviaturaList []*model.Abreviatura
	var abreviaturaObj *model.Abreviatura = &model.Abreviatura{}

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

	return abreviaturaList, nil
}
func (ar *AbreviaturaRepository) GetAbreviaturaByIdProducao(idProducao int64) ([]*model.Abreviatura, error) {

	query := "SELECT a.idAbreviaturaPessoa,a.idPessoa,a.Abreviatura FROM Abreviatura a " +
		"INERN JOIN Producao p " +
		"ON p.idProducao = c.idProducao " +
		"INNER JOIN Coautor c " +
		"ON c.idAbreviatura = a.idAbreviatura " +
		"WHERE p.idProducao = ?"

	rows, err := ar.Connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var abreviaturaList []*model.Abreviatura
	var abreviaturaObj *model.Abreviatura = &model.Abreviatura{}

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

	return abreviaturaList, nil

}

func (ar *AbreviaturaRepository) CreateAbreviatura(abreviatura *model.Abreviatura) error {

	query, err := ar.Connection.Prepare("INSERT INTO Abreviatura (idPessoa,Abreviatura) VALUES (?,?)")

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
