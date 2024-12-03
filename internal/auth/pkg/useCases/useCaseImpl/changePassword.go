package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"log"
	"net/http"
	"strings"
)

func (a *Auth) ChangePassword(payload authDto.ResetRequest) httpresponse.ApiResponse {
	exists, err := a.Repository.Impl.FindByEmail(strings.ToLower(payload.Email), "users", a.Db)
	if err != nil {
		log.Println("Failed to User findByEmail: Email:", payload.Email, ", Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}
	if !exists {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email doesn't exist", nil)
	}

	hashed, err := hash.HashPassword(payload.NewPassword)
	if err != nil {
		log.Println("Failed to hash password ", "Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	//llamar a Impl
	_, err = a.Repository.Impl.ResetPassword(strings.ToLower(payload.Email), hashed, a.Db)
	if err != nil {
		log.Println("Failed to change password ", "Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	return *httpresponse.NewApiError(http.StatusCreated, "password succesfully changed", nil)
}
