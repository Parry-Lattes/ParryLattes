package repository

import (
	"database/sql"
	"fmt"
	"parry_end/model"
)

type CurriculoRepository struct {
	connection *sql.DB
}

func NewCurriculoRepository(connection *sql.DB) CurriculoRepository {
	return CurriculoRepository{
		connection: connection,
	}
}

func (cr *CurriculoRepository) GetCurriculos() ([]model.Curriculo, error) {
	querry := "SELECT idLattes,UltimaAtualizacao FROM Currculo"
	rows, err := cr.connection.Query(querry)

	if err != nil {
		return []model.Curriculo{}, err
	}

	var curriculoList []model.Curriculo
	var curriculoObj model.Curriculo

	for rows.Next() {
		err = rows.Scan(
			&curriculoObj.IdLattes,
			&curriculoObj.UltimaAtualizacao,
		)

		if err != nil {
			return []model.Curriculo{}, err
		}

		curriculoList = append(curriculoList, curriculoObj)

	}

	rows.Close()

	return curriculoList, nil
}

func (cr *CurriculoRepository) GetCurriculoById(idCurrculo int) (*model.Curriculo, error) {
	query, err := cr.connection.Prepare("SELECT c.* " +
		"FROM Currculo c " +
		"WHERE c.idCurrculo = ?")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var curriculo model.Curriculo
	err = query.QueryRow(idCurrculo).Scan(
		&curriculo.IdLattes,
		&curriculo.UltimaAtualizacao,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &curriculo, nil

}
