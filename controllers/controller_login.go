package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"parry_end/model"
	. "parry_end/model"
	. "parry_end/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type ControllerLogin struct {
	loginUsecase *LoginUsecase
}

func NewControllerLogin(loginUse *LoginUsecase) ControllerLogin {
	return ControllerLogin{
		loginUsecase: loginUse,
	}
}

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("Error generating token")
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

func (cl *ControllerLogin) LoginUser(e echo.Context) error {
	var login Login
	err := e.Bind(&login)
	if err != nil {
		response := Response{
			Message: "Body mal-formed!",
		}
		return e.JSON(http.StatusBadRequest, response)
	}

	sessionToken := generateToken(32)
	sessionCookie := &http.Cookie{
		Name:  "session_cookie",
		Value: sessionToken,
		// 30 minutos de duração
		MaxAge:   30 * int(time.Minute/time.Second),
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	e.SetCookie(sessionCookie)
	e.Set("session_token", sessionToken)

	csrfToken := generateToken(32)
	csrfCookie := &http.Cookie{
		Name:     "csrf_cookie",
		Value:    csrfToken,
		MaxAge:   30 * int(time.Minute/time.Second),
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: false,
		SameSite: http.SameSiteStrictMode,
	}
	e.SetCookie(csrfCookie)
	e.Set("csrf_token", csrfToken)

	sessao := model.Sessao{
		IdLogin:     login.IdLogin,
		TokenSessao: sessionToken,
		TokenCSRF:   csrfToken,
	}

	err = cl.loginUsecase.LogUserIn(&login, &sessao)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	return e.NoContent(http.StatusOK)
}
