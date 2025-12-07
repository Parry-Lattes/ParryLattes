package middlewares

import (
	"net/http"
	"parry_end/model"
	"parry_end/usecase"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	loginUsecase *usecase.LoginUsecase
}

func NewAuthMiddleware(loginUse *usecase.LoginUsecase) AuthMiddleware {
	return AuthMiddleware{
		loginUsecase: loginUse,
	}
}

func (auth *AuthMiddleware) CheckIfCSRFTokenExists(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfHeader := c.Request().Header.Get("X-CSRF-Token")
		if csrfHeader == "" {
			res := model.Response{
				Message: "Token CSRF Não foi informado, por favor verifique os headers.",
			}
			return c.JSON(http.StatusUnauthorized, res)
		}

		return next(c)
	}
}

func (auth *AuthMiddleware) CheckIfSessionIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionCookie, err := c.Cookie("session_cookie")
		if err != nil {
			response := model.Response{
				Message: "O cookie de sessão não foi enviado, verifique os headers da requisição.",
			}
			return c.JSON(http.StatusBadRequest, response)
		}

		if sessionCookie == nil {
			response := model.Response{
				Message: "Algum dos tokens de autenticação não foi fornecido, por favor verifique se realizou o login.",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		sessao := model.Sessao{
			TokenSessao: sessionCookie.Value,
			// Esse token sempre vai existir dado que o outro middleware funcionou.
			TokenCSRF: c.Request().Header.Get("X-CSRF-Token"),
		}
		loggedIn, err := auth.loginUsecase.CheckIfIsLoggedIn(&sessao)
		if err != nil {
			response := model.Response{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, response)
		}

		if !loggedIn {
			response := model.Response{
				Message: "O usuário não está logado, por favor realize a autenticação.",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		return next(c)
	}
}
