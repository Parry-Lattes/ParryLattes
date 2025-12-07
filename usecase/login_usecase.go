package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"parry_end/model"
	. "parry_end/repository"
)

var (
	ErrNoLoginFound      = errors.New("O login não está cadastrado no sistema!")
	ErrIncorrectPassword = errors.New("A senha enviada está incorreta.")
	ErrAlreadyExists     = errors.New("Sessão já existe no banco de dados!")
)

type LoginUsecase struct {
	loginRepository  *LoginRepository
	sessaoRepository *SessaoRepository
}

func NewLoginUsecase(
	loginRepo *LoginRepository,
	sessaoRepo *SessaoRepository,
) LoginUsecase {
	return LoginUsecase{
		loginRepository:  loginRepo,
		sessaoRepository: sessaoRepo,
	}
}

func (lu *LoginUsecase) CheckIfIsLoggedIn(sessao *model.Sessao) (bool, error) {
	fmt.Println("Verificando se a sessão Existe", sessao)
	exists, err := lu.sessaoRepository.SessaoExists(sessao)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (lu *LoginUsecase) LogUserIn(
	login *model.Login,
	sessao *model.Sessao,
) error {
	// exists, err := lu.sessaoRepository.SessaoExists(sessao)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// if exists {
	// 	fmt.Println("Sessão já existe no banco de dados!")
	// 	return ErrAlreadyExists
	// }
	//

	fmt.Println("Pegando Login", login)

	dbLogin, err := lu.loginRepository.GetLogin(login)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNoLoginFound
		}

		fmt.Println(err)
		return err
	}

	if login.Senha != dbLogin.Senha {
		return ErrIncorrectPassword
	}

	novaSessao := model.Sessao{
		IdLogin:     dbLogin.IdLogin,
		TokenSessao: sessao.TokenSessao,
		TokenCSRF:   sessao.TokenCSRF,
	}

	fmt.Println("Pegando Sessao por login:", dbLogin)
	// Deletar a sessão já existente no banco se existir.
	sessaoExistente, err := lu.sessaoRepository.GetSessaoByLogin(dbLogin)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	if sessaoExistente != nil {

		fmt.Println("Deletando Sessao por tokens", sessaoExistente)
		err = lu.sessaoRepository.DeleteSessaoByTokens(sessaoExistente)
		if err != nil {
			return err
		}
	}

	fmt.Println("RegistrarSessao",novaSessao)
	// E aí cadastrar a nova sessão no banco
	err = lu.sessaoRepository.RegisterSessao(&novaSessao)
	if err != nil {
		return err
	}

	return nil
}
