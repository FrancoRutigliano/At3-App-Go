package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"log"
	"net/http"
)

func (a *Auth) Login(payload authDto.LoginRequest) httpresponse.ApiResponse {
	user, err := a.Repository.Impl.GetUser(payload, a.Db)
	if err != nil {
		if err.Error() == "not_found" {
			return *httpresponse.NewApiError(http.StatusBadRequest, "user not found", nil)
		}
		log.Println("Failed to get user: Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	if !hash.ComparePasswords(user.Password, []byte(payload.Password)) {
		log.Println("Failed compare password: Error:", err)
		return *httpresponse.NewApiError(http.StatusBadRequest, "invalid credentials", nil)
	}

	token, err := a.JwtService.GenerateTokenRegister(map[string]interface{}{
		"uuid": user.Id,
		"role": user.Role,
	})
	if err != nil {
		log.Println("Failed to generate token: Error:", err)
		return *httpresponse.NewApiError(http.StatusInternalServerError, "An unexpected error occurred", nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, "succesfully login", token)
}
