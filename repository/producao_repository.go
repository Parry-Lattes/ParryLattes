package repository

import (
	"database/sql"
	"fmt"
	"parry_end/model"
)

type ProducaoRepository struct {
	Repository
}

func NewProducaoRepository(connection *sql.DB) ProducaoRepository {
	return ProducaoRepository{
		Repository: Repository{
			Connection: connection,
		},
	}
}
func (pr *ProducaoRepository) GetProducoes() (*[]model.Producao, error) {

	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"INNER JOIN TipoDeProducao tp " +
		"ON p.idTipo = tp.idTipoDeProducao"

	rows, err := pr.Connection.Query(query)

	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return &[]model.Producao{}, err
	}

	var producaoList []model.Producao
	var producaoObj model.Producao

	for rows.Next() {
		err = rows.Scan(
			&producaoObj.IdProducao,
			&producaoObj.Autor,
			&producaoObj.Titulo,
			&producaoObj.Descricao,
			&producaoObj.Link,
			&producaoObj.DataDePublicacao,
			&producaoObj.TipoS,
			&producaoObj.Hash,
		)

		if err != nil {
			return &[]model.Producao{}, err
		}

		producaoList = append(producaoList, producaoObj)
	}

	rows.Close()
	return &producaoList, nil
}

func (pr *ProducaoRepository) GetProducaoByIdLattes(curriculo *model.Curriculo) (*[]model.Producao, error) {
	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"INNER JOIN CurriculoProducao cp " +
		"ON cp.idProducao = p.idProducao " +
		"INNER JOIN Curriculo c " +
		"ON c.idCurriculo = cp.idCurriculo " +
		"INNER JOIN TipoDeProducao tp " +
		"ON tp. idTipoDeProducao = p.idTipo " +
		"WHERE c.idPessoa = ? " +
		"ORDER BY p.DataDePublicacao DESC"

	rows, err := pr.Connection.Query(query, curriculo.IdPessoa)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var producaoList []model.Producao
	var producaoObj model.Producao

	for rows.Next() {

		err = rows.Scan(
			&producaoObj.IdProducao,
			&producaoObj.Autor,
			&producaoObj.Titulo,
			&producaoObj.Descricao,
			&producaoObj.Link,
			&producaoObj.DataDePublicacao,
			&producaoObj.TipoS,
			&producaoObj.Hash,
		)
		if err != nil {
			fmt.Println(err)
			return &[]model.Producao{}, err
		}

		producaoList = append(producaoList, producaoObj)
	}

	rows.Close()

	return &producaoList, nil
}

func (pr *ProducaoRepository) CreateProducao(producao *model.Producao, curriculo *model.Curriculo) (*model.Producao, error) {

	query, err := pr.Connection.Prepare("INSERT INTO Producao (Titulo, idTipo, Descricao, DataDePublicacao, Link, Autor, Hash) " +
		"VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return &model.Producao{}, err
	}

	defer query.Close()

	result, err := query.Exec(
		producao.Titulo,
		producao.TipoId,
		producao.Descricao,
		producao.DataDePublicacao,
		producao.Link,
		producao.Autor,
		producao.Hash,
	)

	if err != nil {
		fmt.Println("Erro ao executar query:", err)
		return &model.Producao{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Errro:", err)
		return &model.Producao{}, err
	}

	producao.IdProducao = id

	return producao, nil
}

func (pr *ProducaoRepository) GetProducaoTypeId(producao *model.Producao) (int, error) {
	query, err := pr.Connection.Prepare("SELECT idTipoDeProducao FROM TipoDeProducao " +
		"WHERE Tipo = ?")

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	var idTipoDeProducao int

	err = query.QueryRow(producao.TipoId).Scan(
		&idTipoDeProducao,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}

		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return idTipoDeProducao, nil

}

func (pr *ProducaoRepository) GetProducaoByHash(producao *model.Producao) (*model.Producao, error) {
	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"WHERE p.Hash = ?"

	rows, err := pr.Connection.Query(query, producao.Hash)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var producaoObj model.Producao

	err = rows.Scan(
		&producaoObj.IdProducao,
		&producaoObj.Autor,
		&producaoObj.Titulo,
		&producaoObj.Descricao,
		&producaoObj.Link,
		&producaoObj.DataDePublicacao,
		&producaoObj.TipoS,
		&producaoObj.Hash,
	)
	if err != nil {
		fmt.Println(err)
		return &model.Producao{}, err
	}

	rows.Close()

	return &producaoObj, nil
}

// func (pr *ProducaoRepository) UpdateProducao(producao *model.Producao, curriculo *model.Curriculo) error {
// 	query, err := pr.Connection.Prepare("UPDATE Producao" +
// 		"SET Autor = ?," +
// 		"Titulo = ?," +
// 		"Descricao = ?," +
// 		"Link = ?," +
// 		"DataDePublicacao = ?," +
// 		"TipoId = ? " +
// 		"WHERE idProducao = ?")

// 	if err != nil {
// 		fmt.Println("Erro ao preparar a query:", err)
// 		return err
// 	}

// 	_, err = query.Exec(
// 		producao.Autor,
// 		producao.Titulo,
// 		producao.Descricao,
// 		producao.Link,
// 		producao.DataDePublicacao,
// 		producao.TipoId,
// 		producao.IdProducao,
// 	)

// 	if err != nil {
// 		fmt.Println("Erro ao executar a query:", err)
// 		return err
// 	}

// 	return nil

// }

// func (pr *ProducaoRepository) IfProducaoExist(producao)
