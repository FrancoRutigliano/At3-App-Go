package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	hash "at3-back/pkg/auth"
	httpresponse "at3-back/pkg/httpResponse"
	"net/http"
)

func (a *Auth) Login(payload authDto.LoginRequest) httpresponse.ApiResponse {
	user, err := a.Repository.Impl.GetUser(payload, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	if !hash.ComparePasswords(user.Password, []byte(payload.Password)) {
		return *httpresponse.NewApiError(http.StatusBadRequest, "invalid credentials", nil)
	}

	token, err := a.JwtService.GenerateTokenRegister(map[string]interface{}{
		"uuid": user.Id,
		"role": user.Role,
	})
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "Oops somenthing went wrong", nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, "succesfully login", token)
}
