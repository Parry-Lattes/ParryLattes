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
			Connection: connection,
		},
	}
}

func (cr *CurriculoRepository) GetCurriculos() ([]*model.Curriculo, error) {
	querry := "SELECT idPessoa,UltimaAtualizacao FROM Curriculo"
	rows, err := cr.Connection.Query(querry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var curriculoList []*model.Curriculo

	for rows.Next() {

		var curriculoObj model.Curriculo = model.Curriculo{}

		err = rows.Scan(
			&curriculoObj.IdPessoa,
			&curriculoObj.UltimaAtualizacao,
		)
		if err != nil {
			return nil, err
		}

		curriculoList = append(curriculoList, &curriculoObj)

	}

	return curriculoList, nil
}

func (cr *CurriculoRepository) GetCurriculoById(
	idPessoa int64,
) (*model.Curriculo, error) {
	fmt.Println(idPessoa)

	query, err := cr.Connection.Prepare(
		"SELECT c.idCurriculo,c.idPessoa, c.UltimaAtualizacao " +
			"FROM Curriculo c " +
			"WHERE c.idPessoa = ?",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var curriculo model.Curriculo

	err = query.QueryRow(idPessoa).Scan(
		&curriculo.IdCurriculo,
		&curriculo.IdPessoa,
		&curriculo.UltimaAtualizacao,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	return &curriculo, nil
}

func (cr *CurriculoRepository) CreateCurriculo(
	curriculo *model.Curriculo,
	pessoa *model.Pessoa,
) (*model.Curriculo, error) {
	query, err := cr.Connection.Prepare(
		"INSERT INTO Curriculo (UltimaAtualizacao,idPessoa)" +
			"VALUES(?,?)",
	)
	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return nil, err
	}

	defer query.Close()

	result, err := query.Exec(
		curriculo.UltimaAtualizacao,
		pessoa.IdPessoa,
	)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao obter último ID:", err)
		return nil, err
	}

	curriculo.IdCurriculo = id

	return curriculo, nil
}

func (cr *CurriculoRepository) UpdateCurriculo(
	curriculo *model.Curriculo,
) error {
	query, err := cr.Connection.Prepare("UPDATE Curriculo " +
		"SET UltimaAtualizacao = ? " +
		"WHERE idPessoa = ?")
	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return err
	}

	defer query.Close()

	result, err := query.Exec(
		curriculo.UltimaAtualizacao,
		curriculo.IdPessoa,
	)

	fmt.Println(result)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil
}

func (cr *CurriculoRepository) LinkCurriculoProducao(
	curriculo *model.Curriculo,
	producao *model.Producao,
) error {
	query, err := cr.Connection.Prepare(
		`INSERT INTO CurriculoProducao (idCurriculo,idProducao)
		VALUES(?,?)`,
	)
	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(
		curriculo.IdCurriculo,
		producao.IdProducao,
	)
	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return err
	}

	return nil
}

func (cr *CurriculoRepository) GetCurriculoId(
	curriculo *model.Curriculo,
) (*int, error) {
	query, err := cr.Connection.Prepare("SELECT c.idCurriculo" +
		"FROM Curriculo c " +
		"WHERE c.idPessoa = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var idCurriculo int

	err = query.QueryRow(curriculo.IdPessoa).Scan(
		&idCurriculo,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	return &idCurriculo, nil
}

func (cu *CurriculoRepository) DeleteCurriculo(idPessoa int64) error {
	query := "DELETE FROM Curriculo WHERE idPessoa = ?"

	result, err := cu.Connection.Exec(query, idPessoa)
	if err != nil {
		fmt.Println("erro ao deletar Curriculo")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("erro ao coletar linhas afetadas")
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("Curriculo não encontrada:")
		return err
	}

	return nil
}

func (cu *CurriculoRepository) UnlinkProducaoCurriculo(
	curriculo *model.Curriculo,
	producao *model.Producao,
) error {
	query := "DELETE FROM CurriculoProducao WHERE idProducao = ?"

	result, err := cu.Connection.Exec(query, producao.IdProducao)
	if err != nil {
		fmt.Println("erro ao deletar CurriculoProducao")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("erro ao coletar linhas afetadas")
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("CurriculoProducao não encontrada:")
		return err
	}

	return nil
}
