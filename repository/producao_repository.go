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

	query := "SELECT p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo " +
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
			&producaoObj.Autor,
			&producaoObj.Titulo,
			&producaoObj.Descricao,
			&producaoObj.Link,
			&producaoObj.DataDePublicacao,
			&producaoObj.Tipo,
		)

		if err != nil {
			return &[]model.Producao{}, err
		}

		producaoList = append(producaoList, producaoObj)
	}

	rows.Close()
	return &producaoList, nil
}

func (pr *ProducaoRepository) GetProducaoById(idProducao int) (*model.Producao, error) {
	query, err := pr.Connection.Prepare("SELECT p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo " +
		"FROM Producao p " +
		"INNER JOIN TipoDeProducao tp " +
		"ON p.idTipo = tp.idTipoDeProducao " +
		"WHERE idProducao = ?")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var producao model.Producao

	err = query.QueryRow(idProducao).Scan(
		&producao.Titulo,
		&producao.Tipo,
		&producao.Descricao,
		&producao.DataDePublicacao,
		&producao.Link,
		&producao.Autor,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &producao, nil
}

// func (pr *ProducaoRepository) CreateProducao(producao *model.Producao, idCurriculo int) error {
// 	query, err := pr.Connection.Prepare("INSERT INTO Producao (Titulo, Tipo, Descricao, DataDePublicacao, Link, Autor) " +
// 		"VALUES (?,?,?,?,?,?)")

// 	if err != nil {
// 		fmt.Println("Erro ao Preparar query:", err)
// 		return err
// 	}

// 	defer query.Close()

// 	result, err := query.Exec(
// 		producao.Titulo,
// 	)
// }
