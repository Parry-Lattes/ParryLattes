package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"parry_end/model"
)

type LoginRepository struct {
	Repository
}

func NewLoginRepository(conn *sql.DB) LoginRepository {
	return LoginRepository{
		Repository: Repository{
			Connection: conn,
		},
	}
}

func (lr *LoginRepository) GetLogin(login *model.Login) (*model.Login, error) {
	query := "SELECT idLogin,Email,Senha FROM Login WHERE Email = ?"
	stmt, err := lr.Connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var dbLogin model.Login
	err = stmt.QueryRow(login.Email).Scan(
		&dbLogin.IdLogin,
		&dbLogin.Email,
		&dbLogin.Senha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Não existe registro desse login!")
			return nil, err
		}

		fmt.Println(err)
		return nil, err
	}

	return &dbLogin, nil
}

func (lr *LoginRepository) CreateLogin(login *model.Login) error {
	return errors.New("Ainda não implementado.")
}
