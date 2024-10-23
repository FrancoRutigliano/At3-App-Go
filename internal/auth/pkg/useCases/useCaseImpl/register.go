package authUseCaseImpl

import (
	authDto "at3-back/internal/auth/pkg/domain/dto"
	httpresponse "at3-back/pkg/httpResponse"
	"net/http"
)

func (a *Auth) Register(payload authDto.RegisterUser) httpresponse.ApiResponse {

	return *httpresponse.NewApiError(http.StatusOK, "user register", nil)
}
