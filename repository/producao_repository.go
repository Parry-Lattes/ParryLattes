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

func (pr *ProducaoRepository) GetProducoes() ([]*model.Producao, error) {
	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"INNER JOIN TipoDeProducao tp " +
		"ON p.idTipo = tp.idTipoDeProducao"

	rows, err := pr.Connection.Query(query)
	if err != nil {
		fmt.Println("Erro ao preparar query:", err)
		return nil, err
	}

	defer rows.Close()

	var producaoList []*model.Producao

	for rows.Next() {

		var producaoObj model.Producao = model.Producao{}

		err = rows.Scan(
			&producaoObj.IdProducao,
			&producaoObj.Autor,
			&producaoObj.Titulo,
			&producaoObj.DataDePublicacao,
			&producaoObj.Tipo,
			&producaoObj.Hash,
		)
		if err != nil {
			return nil, err
		}

		producaoList = append(producaoList, &producaoObj)
	}

	return producaoList, nil
}

func (pr *ProducaoRepository) GetProducaoByIdLattes(
	curriculo *model.Curriculo,
) ([]*model.Producao, error) {
	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"INNER JOIN CurriculoProducao cp " +
		"ON cp.idProducao = p.idProducao " +
		"INNER JOIN Curriculo c " +
		"ON c.idCurriculo = cp.idCurriculo " +
		"INNER JOIN TipoDeProducao tp " +
		"ON tp.idTipoDeProducao = p.idTipo " +
		"WHERE c.idPessoa = ? " +
		"ORDER BY p.DataDePublicacao DESC"

	rows, err := pr.Connection.Query(query, curriculo.IdPessoa)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var producaoList []*model.Producao

	for rows.Next() {

		var producaoObj model.Producao = model.Producao{}

		err = rows.Scan(
			&producaoObj.IdProducao,
			&producaoObj.Autor,
			&producaoObj.Titulo,
			&producaoObj.DataDePublicacao,
			&producaoObj.Tipo,
			&producaoObj.Hash,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		producaoList = append(producaoList, &producaoObj)
	}

	return producaoList, nil
}

func (pr *ProducaoRepository) CreateProducao(
	producao *model.Producao,
	curriculo *model.Curriculo,
) (*model.Producao, error) {
	query, err := pr.Connection.Prepare(
		"INSERT INTO Producao (Titulo, idTipo, DataDePublicacao, Autor, Hash) " +
			"VALUES (?,?,?,?,?)",
	)
	if err != nil {
		fmt.Println("Erro ao Preparar query:", err)
		return nil, err
	}

	fmt.Println(producao.DataDePublicacao)

	defer query.Close()

	result, err := query.Exec(
		producao.Titulo,
		producao.TipoId,
		producao.DataDePublicacao,
		producao.Autor,
		producao.Hash,
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

	producao.IdProducao = id

	return producao, nil
}

func (pr *ProducaoRepository) GetProducaoTypeId(
	producao *model.Producao,
) (int, error) {
	query, err := pr.Connection.Prepare(
		"SELECT idTipoDeProducao FROM TipoDeProducao " +
			"WHERE Tipo = ?",
	)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	defer query.Close()

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

	return idTipoDeProducao, nil
}

func (pr *ProducaoRepository) GetProducaoByHash(
	producao *model.Producao,
) (*model.Producao, error) {
	query := "SELECT p.idProducao,p.Autor,p.Titulo,p.DataDePublicacao,tp.Tipo,p.Hash " +
		"FROM Producao p " +
		"WHERE p.Hash = ?"

	rows, err := pr.Connection.Query(query, producao.Hash)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	err = rows.Scan(
		&producao.IdProducao,
		&producao.Autor,
		&producao.Titulo,
		&producao.DataDePublicacao,
		&producao.Tipo,
		&producao.Hash,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return producao, nil
}

func (pr *ProducaoRepository) DeleteProducaoByIdCurriculo(
	idCurriculo int64,
) error {
	query := "DELETE c.*,p.* FROM Producao p " +
		"INNER JOIN CurriculoProducao c " +
		"ON c.idProducao = p.idProducao " +
		"WHERE c.idCurriculo = ?"

	result, err := pr.Connection.Exec(query, idCurriculo)
	if err != nil {
		fmt.Println("erro ao deletar Producoes")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("erro ao coletar linhas afetadas")
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("Producao não encontrada:")
		return err
	}

	return nil
}

//
// func (pr *ProducaoRepository) DeleteProducaoBydHash(hash int64) error {
// 	query := "DELETE FROM Producao WHERE Hash = ?"
//
// 	result, err := pr.Connection.Exec(query, hash)
// 	if err != nil {
// 		fmt.Println("erro ao deletar producao")
// 		return err
// 	}
//
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		fmt.Println("erro ao coletar linhas afetadas")
// 		return err
// 	}
//
// 	if rowsAffected == 0 {
// 		fmt.Println("producao não encontrada:")
// 		return err
// 	}
//
// 	return nil
//
// }

func (pr *ProducaoRepository) GetCoautoresById(
	IdProducao *int64,
) ([]*model.Abreviatura, error) {
	query := "SELECT a.idAbreviaturaPessoa,a.idPessoa,a.Abreviatura WHERE from Abreviatura a " +
		"INNER JOIN Producao p " +
		"ON p.idProducao = a.idProducao " +
		"INNER JOIN Coautor c " +
		"ON c.idAbreviatura = a.idAbreviatura " +
		"WHERE p.idProducao = ?"

	rows, err := pr.Connection.Query(query, IdProducao)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var coautorList []*model.Abreviatura

	for rows.Next() {

		var coautorObj model.Abreviatura = model.Abreviatura{}

		err = rows.Scan(
			&coautorObj.IdAbreviatura,
			&coautorObj.IdPessoa,
			&coautorObj.Abreviatura,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			fmt.Println(err)
			return nil, err
		}

		coautorList = append(coautorList, &coautorObj)
	}

	return coautorList, nil
}

// func (pr *ProducaoRepository) CreateAbreviatura(
// 	abreviatura *model.Abreviatura,
// ) (*model.Abreviatura, error) {
// 	query, err := pr.Connection.Prepare(
// 		"INSERT INTO Abreviatura (idPessoa,Abreviatura)" +
// 			"VALUES(?,?)",
// 	)
// 	if err != nil {
// 		fmt.Println("Erro ao Preparar query:", err)
// 		return nil, err
// 	}
//
// 	defer query.Close()
//
// 	result, err := query.Exec(
// 		abreviatura.IdPessoa,
// 		abreviatura.IdAbreviatura,
// 	)
// 	if err != nil {
// 		fmt.Println("Erro ao executar query:", err)
// 		return nil, err
// 	}
//
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		fmt.Println("Erro ao obter último ID:", err)
// 		return nil, err
// 	}
//
// 	abreviatura.IdAbreviatura = id
//
// 	return abreviatura, nil
// }

// func (pr *ProducaoRepository) GetProducaoBibliograficaCount(
//
//	ano int,
//
//	) (*int64, error) {
//		query, err := pr.Connection.Prepare("SELECT COUNT(*) FROM Producao p " +
//			"INNER JOIN TipoDeProducao tp " +
//			"ON p.idTipo = tp.idTipoDeProducao " +
//			"WHERE tp.idTipoDeProducao = 1 " +
//			"AND p.DataDePublicacao = ?")
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//
//		defer query.Close()
//
//		var contagem *int64
//
//		err = query.QueryRow(ano).Scan(contagem)
//		if err != nil {
//			return nil, err
//		}
//
//		return contagem, nil
//	}
//
// func (pr *ProducaoRepository) GetProducaoTecnicaCountByAno(
//
//	ano int,
//
//	) (*int64, error) {
//		query, err := pr.Connection.Prepare("SELECT COUNT(*) FROM Producao p " +
//			"INNER JOIN TipoDeProducao tp " +
//			"ON p.idTipo = tp.idTipoDeProducao " +
//			"WHERE tp.idTipoDeProducao = 2 " +
//			"AND p.DataDePublicacao = ?")
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//
//		defer query.Close()
//
//		var contagem *int64
//
//		err = query.QueryRow(ano).Scan(contagem)
//		if err != nil {
//			return nil, err
//		}
//
//		return contagem, nil
//	}
//
// func (pr *ProducaoRepository) GetProducaoPatenteCountByAno(
//
//	ano int,
//
//	) (*int64, error) {
//		query, err := pr.Connection.Prepare("SELECT COUNT(*) FROM Producao p " +
//			"INNER JOIN TipoDeProducao tp " +
//			"ON p.idTipo = tp.idTipoDeProducao " +
//			"WHERE tp.idTipoDeProducao = 3 " +
//			"AND p.DataDePublicacao = ?")
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//
//		defer query.Close()
//
//		var contagem *int64
//
//		err = query.QueryRow(ano).Scan(contagem)
//		if err != nil {
//			return nil, err
//		}
//
//		return contagem, nil
//	}
//
// func (pr *ProducaoRepository) GetProducaoOutroCountByANo(
//
//	ano int,
//
//	) (*int64, error) {
//		query, err := pr.Connection.Prepare("SELECT COUNT(*) FROM Producao p " +
//			"INNER JOIN TipoDeProducao tp " +
//			"ON p.idTipo = tp.idTipoDeProducao " +
//			"WHERE tp.idTipoDeProducao = 4 " +
//			"AND p.DataDePublicacao = ?")
//		if err != nil {
//			fmt.Println(err)
//			return nil, err
//		}
//
//		defer query.Close()
//
//		var contagem *int64
//
//		err = query.QueryRow(ano).Scan(contagem)
//		if err != nil {
//			return nil, err
//		}
//
//		return contagem, nil
//	}
//
// func (pr *ProducaoRepository) GetProducoesYearAndType() ([]*model.TupleAnoTipo, error) {
// }
func (pr *ProducaoRepository) GetProducaoCount() (*int64, error) {
	query, err := pr.Connection.Prepare("SELECT COUNT(*) FROM Producao p")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var contagem int64

	err = query.QueryRow().Scan(&contagem)
	if err != nil {
		return nil, err
	}

	return &contagem, nil
}

func (pr *ProducaoRepository) GetProducoesGroypByAnoTipo() (map[int]*model.RelatorioAno, error) {
	query := "SELECT p.DataDePublicacao as ano, p.idTipo, COUNT(*) as quantidade " +
		"FROM Producao p " +
		"WHERE p.DataDePublicacao IS NOT NULL " +
		"GROUP BY p.DataDePublicacao, p.idTipo " +
		"ORDER BY ano DESC"

	rows, err := pr.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	relatorios := make(map[int]*model.RelatorioAno)

	fmt.Println("sexo")

	for rows.Next() {
		var ano int
		var tipo int
		var quantidade int64

		err = rows.Scan(&ano, &tipo, &quantidade)
		if err != nil {
			return nil, err
		}

		if _, exists := relatorios[ano]; !exists {
			relatorios[ano] = &model.RelatorioAno{}
		}

		relatorioAno := relatorios[ano]

		// Distribui pelos tipos
		switch tipo {
		case 1:
			relatorioAno.ProducoesBibliograficas = &quantidade
		case 2:
			relatorioAno.ProducoesTecnicas = &quantidade
		case 3:
			relatorioAno.ProducoesPatente = &quantidade
		case 4:
			relatorioAno.ProducoesOutro = &quantidade
		}

		// Calcula total
		if relatorioAno.ProducoesTotal == nil {
			relatorioAno.ProducoesTotal = new(int64)
		}
		*relatorioAno.ProducoesTotal += quantidade
	}

	return relatorios, nil
}

func (pr *ProducaoRepository) GetProducaoCountByIdLattes(
	idLattes string,
) (*int64, error) {
	query := "SELECT COUNT(*) FROM Producao p " +
		"INNER JOIN CurriculoProducao cp " +
		"ON cp.idProducao = p.idProducao " +
		"INNER JOIN Curriculo c " +
		"ON c.idCurriculo = cp.idCurriculo " +
		"INNER JOIN Pessoa ps " +
		"ON ps.idPessoa = c.idPessoa " +
		"WHERE ps.idLattes = ?"

	rows, err := pr.Connection.Query(query, idLattes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var contagem *int64

	err = rows.Scan(
		contagem,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return contagem, nil
}

// func (pr *ProducaoRepository) UpdateProducao(producao *model.Producao, curriculo *model.Curriculo) error {
// 	query, err := pr.Connection.Prepare("UPDATE Producao" +
// 		"SET Autor = ?," +
// 		"Titulo = ?," +
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
