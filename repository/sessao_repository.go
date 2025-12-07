package repository

import (
	"database/sql"
	"fmt"

	"parry_end/model"
)

type SessaoRepository struct {
	Repository
}

func NewSessaoRepository(conn *sql.DB) SessaoRepository {
	return SessaoRepository{
		Repository{
			Connection: conn,
		},
	}
}

func (sr *SessaoRepository) SessaoExists(sessao *model.Sessao) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM Sessao WHERE TokenSessao = ? AND TokenCsrf = ? LIMIT 1)"
	stmt, err := sr.Connection.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(sessao.TokenSessao, sessao.TokenCSRF).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (sr *SessaoRepository) GetSessaoByLogin(
	login *model.Login,
) (*model.Sessao, error) {
	query := "SELECT idSessao, TokenSessao, TokenCsrf FROM Sessao WHERE idLogin = ?"
	stmt, err := sr.Connection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var sessao model.Sessao
	err = stmt.QueryRow(login.IdLogin).
		Scan(&sessao.IdSessao, &sessao.TokenSessao, &sessao.TokenCSRF)
	if err != nil {
		return nil, err
	}

	return &sessao, nil
}

func (sr *SessaoRepository) RegisterSessao(sessao *model.Sessao) error {
	query := "INSERT INTO Sessao (idLogin, TokenSessao, TokenCsrf) VALUES (?, ?, ?)"

	stmt, err := sr.Connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&sessao.IdLogin, sessao.TokenSessao, sessao.TokenCSRF)
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessaoRepository) DeleteSessaoByTokens(sessao *model.Sessao) error {
	query := "DELETE FROM Sessao WHERE TokenSessao = ? AND TokenCsrf = ?"
	stmt, err := sr.Connection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessao.TokenSessao, sessao.TokenCSRF)
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessaoRepository) DeleteSessaoByLogin(login *model.Login) error {
	query := "DELETE FROM Sessao WHERE idLogin = ?"
	stmt, err := sr.Connection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(login.IdLogin)
	if err != nil {
		return err
	}

	return nil
}
