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

func (pr *ProducaoRepository) GetProducaoById(idLattes int) (*[]model.Producao, error) {
	query := ("SELECT p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo " +
		"FROM Producao p " +
		"INNER JOIN CurriculoProducao cp " +
		"ON cp.idProducao = p.idProducao " +
		"INNER JOIN Curriculo c " +
		"ON c.idCurriculo = cp.idCurriculo " +
		"INNER JOIN TipoDeProducao tp " +
		"ON tp. idTipoDeProducao = p.idTipo " +
		"WHERE c.idLattes = ? " +
		"ORDER BY p.DataDePublicacao DESC")

	rows, err := pr.Connection.Query(query, idLattes)

	if err != nil {
		fmt.Println(err)
		return nil, err
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
			fmt.Println(err)
			return &[]model.Producao{}, err
		}

		producaoList = append(producaoList, producaoObj)
	}

	rows.Close()

	return &producaoList, nil
}

// func (pr *ProducaoRepository) CreateProducao(producao *model.Producao, idCurriculo int, idTipo int) error {

// 	query, err := pr.Connection.Prepare("INSERT INTO Producao (Titulo, idTipo, Descricao, DataDePublicacao, Link, Autor) " +
// 		"VALUES (?,?,?,?,?,?)")

// 	if err != nil {
// 		fmt.Println("Erro ao Preparar query:", err)
// 		return err
// 	}

// 	defer query.Close()

// 	result, err := query.Exec(
// 		producao.Titulo,
// 		idTipo,
// 		producao.Descricao,
// 		producao.DataDePublicacao,
// 		producao.Link,
// 		producao.Autor,
// 	)

// 	if err != nil {
// 		fmt.Println("Erro ao executar query:", err)
// 		return err
// 	}

// 	id, err := result.LastInsertId() // Necessário dar uso para result

// 	if err != nil {
// 		fmt.Println("Errro:", err)
// 		return err
// 	}

// 	query, err = pr.Connection.Prepare("INSERT INTO CurriculoProducao (idCurriculo,idProducao) " +
// 		"VALUES = (?,?)")

// 	result, err = query.Exec(
// 		idCurriculo,
// 		id,
// 	)

// 	if err != nil {
// 		fmt.Println("Erro ao executar query:", err)
// 		return err
// 	}

// 	fmt.Println(result)

// 	return nil
// }
