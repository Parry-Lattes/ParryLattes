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

	query := "SELECT p.Autor,p.Titulo,p.Descricao,p.Link,p.DataDePublicacao,tp.Tipo FROM Producao p INNER JOIN TipoDeProducao tp ON p.idTipo = tp.idTipoDeProducao"

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
