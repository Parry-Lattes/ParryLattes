package repository

import (
	"database/sql"
	"fmt"
	"parry_end/model"
)

type CurriculoRepository struct {
	Repository
}

func NewCurriculoRepository(connection *sql.DB) CurriculoRepository {
	return CurriculoRepository{
		Repository: Repository{
			Connection: connection},
	}
}

func (cr *CurriculoRepository) GetCurriculos() (*[]model.Curriculo, error) {
	querry := "SELECT idLattes,UltimaAtualizacao FROM Curriculo"
	rows, err := cr.Connection.Query(querry)

	if err != nil {
		return &[]model.Curriculo{}, err
	}

	var curriculoList []model.Curriculo
	var curriculoObj model.Curriculo

	for rows.Next() {
		err = rows.Scan(
			&curriculoObj.IdLattes,
			&curriculoObj.UltimaAtualizacao,
		)

		if err != nil {
			return &[]model.Curriculo{}, err
		}

		curriculoList = append(curriculoList, curriculoObj)

	}

	rows.Close()

	return &curriculoList, nil
}

func (cr *CurriculoRepository) GetCurriculoById(idLattes int) (*model.Curriculo, error) {
	query, err := cr.Connection.Prepare("SELECT c.idLattes, c.UltimaAtualizacao " +
		"FROM Curriculo c " +
		"WHERE c.idLattes = ?")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var curriculo model.Curriculo
	err = query.QueryRow(idLattes).Scan(
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

func (cr *CurriculoRepository) CreateCurriculo(curriculo *model.Curriculo, idPessoa int) (int, error) {
	query, err := cr.Connection.Prepare(`INSERT INTO Curriculo (idLattes,UltimaAtualizacao,idPessoa)
		VALUES(?,?,²?)`)

	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return 0, err
	}

	defer query.Close()

	result, err := query.Exec(
		curriculo.IdLattes,
		curriculo.UltimaAtualizacao,
		idPessoa,
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

func (cr *CurriculoRepository) UpdateCurriculo(curriculo *model.Curriculo) error {
	query, err := cr.Connection.Prepare("UPDATE Curriculo " +
		"SET UltimaAtualizacao = ? " +
		"WHERE idLattes = ?")

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	result, err := query.Exec(
		curriculo.UltimaAtualizacao,
		curriculo.IdLattes,
	)

	fmt.Println(result)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil

}
